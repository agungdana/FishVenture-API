package ptime_test

import (
	"fmt"
	"testing"

	"github.com/e-fish/api/pkg/common/helper/ptime"
)

func TestTime(t *testing.T) {
	ptime.SetDefaultTimeToUTC()

	fmt.Printf("ptime.Today(): %v\n", ptime.Today())
}
