package migrations_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("i:%v = %v\n", i, uuid.New())
	}
}
