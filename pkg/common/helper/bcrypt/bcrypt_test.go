package bcrypt_test

import (
	"testing"

	"github.com/e-fish/api/pkg/common/helper/bcrypt"
	"github.com/stretchr/testify/assert"
)

func Test_Compare(t *testing.T) {
	hashedPass := "$2a$10$r6huirn1laq6UXBVu6ga9.sHca6sr6tQl3Tiq9LB6/6LMpR37XEGu"
	pass := "password"
	// hashedPassNew, _ := bcrypt.HashPassowrd(pass)
	err := bcrypt.ComparePassword(pass, hashedPass)
	assert.Empty(t, err, "errors")
}
