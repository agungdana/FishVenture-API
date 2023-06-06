package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func newFireBaseMessaging(app *firebase.App, ctx context.Context) (Messaging, error) {
	msg, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseMessaging{
		msg: msg,
	}, nil
}

type firebaseMessaging struct {
	msg *messaging.Client
}

func (f *firebaseMessaging) SendMessage(ctx context.Context, data FirebaseMessageData) {
	f.msg.Send(ctx, &messaging.Message{
		Data: data.Data,
		Notification: &messaging.Notification{
			Title: data.Title,
			Body:  data.Body,
		},
		Android: &messaging.AndroidConfig{
			Data: data.Data,
			Notification: &messaging.AndroidNotification{
				Title: data.Title,
				Body:  data.Body,
			},
		},
		Topic: data.Topic,
	})
}
