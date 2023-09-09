package internal

import (
	"fmt"
	"time"
)

type Scheduler struct {
	hour  int
	Timer *time.Timer
}

func NewScheduler(hour int) *Scheduler {
	return &Scheduler{
		hour: hour,
	}
}

func (s *Scheduler) SetSchedule() {
	today := time.Now()
	newTimer := time.Date(today.Year(), today.Month(), today.Day(), s.hour, 0, 0, 0, today.Location())

	fmt.Printf("today: %v\n", today)
	fmt.Printf("newTimer: %v\n", newTimer)

	if newTimer.Before(today) {
		newTimer = newTimer.AddDate(0, 0, 1)
	}

	diff := newTimer.Sub(today)
	fmt.Printf("diff: %v\n", diff)
	if s.Timer == nil {
		s.Timer = time.NewTimer(diff)
		return
	}

	s.Timer.Reset(diff)

}
