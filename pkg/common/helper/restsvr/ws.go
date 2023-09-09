package restsvr

import (
	"context"
	"encoding/json"
	"strings"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsManager struct {
	mut  sync.Mutex
	subs event.EventPubsub
	conf config.AppConfig

	listConnection map[uuid.UUID]*conn
}

type conn struct {
	conn *websocket.Conn
	msg  chan []byte
	done chan struct{}
}

func (c *conn) Close() error {
	return c.conn.Close()
}

func (c *conn) ReadMessage() {
	for {
		_, wsMsg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("ws close error: %v", err)
			}
			c.done <- struct{}{}
			break
		}

		if wsMsg != nil {
			c.msg <- wsMsg
		}
	}
}

func NewWsManager(conf config.AppConfig, event event.Event) *WsManager {
	return &WsManager{
		mut:            sync.Mutex{},
		subs:           event.NewEventPubsub(),
		conf:           conf,
		listConnection: make(map[uuid.UUID]*conn),
	}
}

func (wm *WsManager) getWsConn(ctx context.Context, wsConn ...*websocket.Conn) *conn {
	var (
		requestID, _ = ctxutil.GetRequestID(ctx)
	)

	if ws, ok := wm.listConnection[requestID]; ok {
		return ws
	}

	if len(wsConn) > 0 {
		newWsConn := &conn{
			conn: wsConn[0],
			msg:  make(chan []byte),
			done: make(chan struct{}),
		}
		wm.listConnection[requestID] = newWsConn
		return newWsConn
	}

	logger.Fatal("####### failed add/found wsConn, need websocket.Conn")
	return nil
}

func (wm *WsManager) WsRun(ctx context.Context, wsConnection *websocket.Conn) error {
	var (
		topic = event.GetTopic(ctx, string(event.WS_REPLY), wm.conf.Debug)
		ws    = wm.getWsConn(ctx, wsConnection)
	)

	id, err := wm.subs.Subscribe(ctx, topic, func(m *event.ClientMessages) {
		data, ok := m.Data.([]byte)
		if !ok {
			dataMarshal, err := json.Marshal(&m.Data)
			if err != nil {
				logger.ErrorWithContext(ctx, "### Failed marshal message data")
				return
			}

			data = dataMarshal
		}

		err := ws.conn.WriteMessage(1, data)
		if err != nil {
			logger.ErrorWithContext(ctx, "### Failed marshal message data")
			return
		}
	})

	if err != nil {
		return err
	}

	go wm.waitMessage(ctx, id)

	return nil
}

type WsMessage struct {
	Topic  string `json:"topic"`
	Action string `json:"action"`
	Data   any    `json:"data"`
}

type WsResponse struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func (wm *WsManager) waitMessage(ctx context.Context, subsID uuid.UUID) {
	ctx = ctxutil.CopyCtxWithoutTimeout(ctx)
	var (
		ws     = wm.getWsConn(ctx)
		conf   = wm.conf
		ctxMap = ctxutil.ToMap(ctx)
	)

	defer wm.wsClose(ctx, subsID)

	go ws.ReadMessage()

forClose:
	for {
		select {
		case data := <-ws.msg:
			if strings.ToUpper(string(data)) == "PING" {
				err := ws.conn.WriteMessage(1, []byte("PONG"))
				if err != nil {
					logger.ErrorWithContext(ctx, "### Failed marshal message data")
				}
				continue
			}

			wsMsg := WsMessage{}
			err := json.Unmarshal(data, &wsMsg)
			if err != nil {
				logger.ErrorWithContext(ctx, "failed json.Unmarshal ws.msg with err [%v]", err)
				continue
			}

			if wsMsg.Topic == "" {
				wsMsg.Topic = string(event.WS_COLLECT)
			}

			topic := event.GetTopic(ctx, wsMsg.Topic, conf.Debug)

			err = wm.subs.Publish(ctx, event.ClientMessages{
				Topic:  topic,
				Action: wsMsg.Action,
				CtxMap: ctxMap,
				Data:   wsMsg.Data,
			})

			if err != nil {
				err = werror.Error{
					Code:    "FailedPublishWsMessages",
					Message: "failed publish messages",
					Details: map[string]any{
						"topic":  topic,
						"action": wsMsg.Action,
						"error":  err,
						"data":   wsMsg.Data,
					},
				}
				logger.ErrorWithContext(ctx, "###Failed publish message ws err: %v", err)
			}
		case <-ws.done:
			break forClose
		}
	}
	close(ws.msg)
	close(ws.done)
}

func (wm *WsManager) wsClose(ctx context.Context, subsID uuid.UUID) {
	var (
		requestID, _ = ctxutil.GetRequestID(ctx)
		ws, ok       = wm.listConnection[requestID]
	)

	wm.mut.Lock()

	if ok {
		ws.Close()
		delete(wm.listConnection, requestID)
	}

	logger.Debug("###ws close with req id [%v]", requestID)
	wm.subs.Close(ctx, subsID)

	wm.mut.Unlock()
}
