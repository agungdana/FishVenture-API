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
	return fromContextString(ctx, USER_ID)
}

func SetUserPayload(ctx context.Context, userID, PondID uuid.UUID, roleID ...uuid.UUID) context.Context {
	ctx = SetUserID(ctx, userID)
	ctx = SetPondID(ctx, PondID)
	ctx = SetRoleID(ctx, roleID...)
	return ctx
}
