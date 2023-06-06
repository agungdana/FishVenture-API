package firebase

import (
	"context"
)

type Firebase interface {
	NewGoogleAuth(ctx context.Context) (GoogleAuth, error)
	NewMessaging(ctx context.Context) (Messaging, error)
}

type GoogleAuth interface {
	Signin(ctx context.Context, idToken string) (sign *Signature, err error)
}

type Messaging interface {
	SendMessage(ctx context.Context, data FirebaseMessageData)
}
