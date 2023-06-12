package firebase

import (
	"context"

	"os"

	firebase "firebase.google.com/go"
	"github.com/e-fish/api/pkg/common/helper/config"
	"google.golang.org/api/option"
)

func NewFirebase(conf config.FirebaseConfig) (Firebase, error) {

	data, err := os.ReadFile(conf.FireBase)
	if err != nil {
		return nil, err
	}
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON(data))
	if err != nil {
		return nil, err
	}

	return &fireBase{
		app: app,
	}, nil
}

type fireBase struct {
	app *firebase.App
}

// NewGoogleAuth implements Firebase
func (f *fireBase) NewGoogleAuth(ctx context.Context) (GoogleAuth, error) {
	return newAuth(ctx, f.app)
}

// NewMessaging implements Firebase
func (f *fireBase) NewMessaging(ctx context.Context) (Messaging, error) {
	return newFireBaseMessaging(f.app, ctx)
}
