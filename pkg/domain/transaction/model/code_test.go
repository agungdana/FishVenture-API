package model_test

import (
	"fmt"
	"testing"

	"github.com/e-fish/api/pkg/domain/transaction/model"
)

func TestCode(t *testing.T) {
	code := model.GenerateCode()
	fmt.Printf("code: %v\n", code)
}
