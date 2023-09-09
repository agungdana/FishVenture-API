package ctxutil

import (
	"context"

	"github.com/google/uuid"
)

type key string

const (
	REQUEST_ID     key = "X-Efish-Request-ID"
	TRANSACTION_ID key = "X-Efish-Transaction-ID"
	USER_ID        key = "X-Efish-User-ID"
	POND_ID        key = "X-Efish-Pond-ID"
	ROLE_ID        key = "X-Efish-Role-ID"
	AppType        key = "X-Efish-App-Type"
)

func fromContextUUID(ctx context.Context, key key) (uuid.UUID, bool) {
	value := ctx.Value(key)
	switch v := value.(type) {
	case string:
		if uid, err := uuid.Parse(v); err == nil {
			if uid == uuid.Nil {
				return uid, false
			}
			return uid, true
		}
		return uuid.UUID{}, false
	case uuid.UUID:
		if v == uuid.Nil {
			return v, false
		}
		return v, true
	default:
		return uuid.UUID{}, false
	}
}

func fromContextString(ctx context.Context, key key) (string, bool) {
	value := ctx.Value(key)
	switch v := value.(type) {
	case string:
		return v, true
	default:
		return "", false
	}
}

func fromContextArrayUUID(ctx context.Context, key key) ([]uuid.UUID, bool) {
	value := ctx.Value(key)
	switch v := value.(type) {
	case string:
		return nil, false
	case []uuid.UUID:
		return v, true
	case []*uuid.UUID:
		uid := []uuid.UUID{}
		for _, v := range v {
			uid = append(uid, *v)
		}
		return uid, true
	default:
		return nil, false
	}
}

func NewTransactionID(ctx context.Context) context.Context {
	return context.WithValue(ctx, TRANSACTION_ID, uuid.New())
}

func NewRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, REQUEST_ID, uuid.New())
}

func NewRequest(ctx context.Context) context.Context {
	ctx = NewTransactionID(ctx)
	ctx = NewRequestID(ctx)
	return ctx
}

func NewRequestWithOutTimeOut(ctx context.Context) context.Context {
	newCtx := context.Background()
	newCtx = NewTransactionID(newCtx)
	newCtx = NewRequestID(newCtx)
	userID, _ := GetUserID(ctx)
	roleID, _ := GetRoleID(ctx)
	pondID, _ := GetPondID(ctx)
	appType, _ := GetUserAppType(ctx)
	newCtx = SetUserID(newCtx, userID)
	newCtx = SetRoleID(newCtx, roleID...)
	newCtx = SetPondID(newCtx, pondID)
	newCtx = SetUserAppType(newCtx, appType)
	return newCtx
}

func GetRequestID(ctx context.Context) (uuid.UUID, bool) {
	return fromContextUUID(ctx, REQUEST_ID)
}

func GetTransactionID(ctx context.Context) (uuid.UUID, bool) {
	return fromContextUUID(ctx, TRANSACTION_ID)
}

func SetUserID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, USER_ID, id)
}

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	return fromContextUUID(ctx, USER_ID)
}

func SetRoleID(ctx context.Context, id ...uuid.UUID) context.Context {
	return context.WithValue(ctx, ROLE_ID, id)
}

func GetRoleID(ctx context.Context) ([]uuid.UUID, bool) {
	return fromContextArrayUUID(ctx, ROLE_ID)
}

func SetPondID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, POND_ID, id)
}

func GetPondID(ctx context.Context) (uuid.UUID, bool) {
	return fromContextUUID(ctx, POND_ID)
}

func SetUserAppType(ctx context.Context, appType string) context.Context {
	return context.WithValue(ctx, AppType, appType)
}

func GetUserAppType(ctx context.Context) (string, bool) {
	return fromContextString(ctx, AppType)
}

func SetUserPayload(ctx context.Context, userID, PondID uuid.UUID, appType string, roleID ...uuid.UUID) context.Context {
	ctx = SetUserID(ctx, userID)
	ctx = SetPondID(ctx, PondID)
	ctx = SetRoleID(ctx, roleID...)
	ctx = SetUserAppType(ctx, appType)
	return ctx
}

func CopyCtxWithoutTimeout(ctx context.Context) context.Context {
	newContext := context.Background()
	requestID, _ := GetRequestID(ctx)
	newContext = context.WithValue(newContext, REQUEST_ID, requestID)
	transactionID, _ := GetTransactionID(ctx)
	newContext = context.WithValue(newContext, TRANSACTION_ID, transactionID)
	userID, _ := GetUserID(ctx)
	newContext = SetUserID(newContext, userID)
	roleID, _ := GetRoleID(ctx)
	newContext = SetRoleID(newContext, roleID...)
	pondID, _ := GetPondID(ctx)
	newContext = SetPondID(newContext, pondID)
	appType, _ := GetUserAppType(ctx)
	newContext = SetUserAppType(newContext, appType)
	return newContext
}

func ToMap(ctx context.Context) map[string]any {
	ctxMap := make(map[string]any)

	if userID, ok := GetUserID(ctx); ok {
		ctxMap[string(USER_ID)] = userID
	}

	if roleID, ok := GetRoleID(ctx); ok {
		ctxMap[string(ROLE_ID)] = roleID
	}

	if pondID, ok := GetPondID(ctx); ok {
		ctxMap[string(POND_ID)] = pondID
	}

	if appType, ok := GetUserAppType(ctx); ok {
		ctxMap[string(AppType)] = appType
	}

	return ctxMap
}

func ToContextUsingMap(ctxMap map[string]any) context.Context {
	ctx := context.Background()

	if userID, ok := ctxMap[string(USER_ID)]; ok {
		ctx = SetUserID(ctx, ToUUID(userID))
	}

	if roleID, ok := ctxMap[string(ROLE_ID)]; ok {
		ctx = SetRoleID(ctx, ToUUID(roleID))
	}

	if pondID, ok := ctxMap[string(POND_ID)]; ok {
		ctx = SetPondID(ctx, ToUUID(pondID))
	}

	if appType, ok := ctxMap[string(AppType)]; ok {
		ctx = SetUserAppType(ctx, ToString(appType))
	}

	return ctx
}

func ToUUID(data any) uuid.UUID {
	switch v := data.(type) {
	case uuid.UUID:
		return v
	case string:
		return uuid.MustParse(v)
	case []byte:
		uid, _ := uuid.FromBytes(v)
		return uid
	}
	return uuid.Nil
}

func ToString(data any) string {
	switch v := data.(type) {
	case uuid.UUID:
		return v.String()
	case string:
		return v
	case []byte:
		return string(v)
	}
	return ""
}
