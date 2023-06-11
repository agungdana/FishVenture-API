package firebase

import (
	"context"

	"encoding/json"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/e-fish/api/pkg/common/helper/logger"
)

func newAuth(ctx context.Context, app *firebase.App) (GoogleAuth, error) {

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &goAuth{authClient: auth}, nil
}

type goAuth struct {
	authClient *auth.Client
}

// Signin implements GoogleAuth
func (g *goAuth) Signin(ctx context.Context, idToken string) (sign *Signature, err error) {
	var (
		emails []string
		email  string
	)

	token, err := g.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		logger.InfoWithContext(ctx, "failed verify token id: %v", err)
		return nil, err
	}

	// fmt.Printf("token: %v\n", token)
	m, ok := token.Firebase.Identities["email"]
	if !ok {
		return nil, errors.New("need email")
	}

	data, _ := json.Marshal(m)

	json.Unmarshal(data, &emails)
	if len(emails) > 0 {
		email = emails[0]
	}

	return &Signature{
		Email: email,
	}, nil

}
