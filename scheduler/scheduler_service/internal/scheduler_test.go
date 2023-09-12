package internal_test

import (
	"fmt"
	"testing"

	"github.com/e-fish/api/scheduler/scheduler_service/internal"
)

func Test_Name(t *testing.T) {
	sc := internal.NewScheduler(41)

	sc.SetSchedule()

	<-sc.Timer.C
	fmt.Println("###############Runnnnnnn")
	sc.SetSchedule()
}
