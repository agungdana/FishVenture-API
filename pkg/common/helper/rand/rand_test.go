package rand_test

import (
	"testing"

	"github.com/e-fish/api/pkg/common/helper/rand"
	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	code := rand.RandCode(10)
	assert.NotEmpty(t, code)
}
