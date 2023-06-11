package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(p string) (string, error) {

	result, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(password, hashPassowrd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassowrd), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
