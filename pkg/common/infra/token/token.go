package token

import (
	"strings"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
)

type Token interface {
	CreateToken(payload Claims) (string, error)
	VerifyToken(token string, payload Claims) error
}

func NewTokenMaker(secret string) (Token, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(strings.NewReader(secret))
	if err != nil {
		return nil, err
	}

	return &tokenMaker{
		maker: paseto.NewV2(),
		pub:   publicKey,
		priv:  privateKey,
	}, nil
}

type tokenMaker struct {
	maker *paseto.V2
	pub   ed25519.PublicKey
	priv  ed25519.PrivateKey
}

// CreateToken implements Token
func (m *tokenMaker) CreateToken(payload Claims) (string, error) {
	return m.maker.Sign(m.priv, payload, nil)
}

// VerifyToken implements Token
func (m *tokenMaker) VerifyToken(token string, payload Claims) error {
	err := m.maker.Verify(token, m.pub, &payload, nil)
	if err != nil {
		return err
	}

	err = payload.Valid()
	if err != nil {
		return err
	}

	return nil
}
