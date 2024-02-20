package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperations struct {
	Ops []string
}

func (c *CountdownOperations) Sleep() {
	c.Ops = append(c.Ops, "sleep")
}

func (c *CountdownOperations) Write(b []byte) (int, error) {
	c.Ops = append(c.Ops, "write")
	return 0, nil
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {

	t.Run("should print countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		expected := `3
2
1
Go!`

		Countdown(buffer, spySleeper)
		got := buffer.String()

		if got != expected {
			t.Errorf("expected %q, but got %q", expected, got)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("sleeper not called enought, expected 3 but got %d", spySleeper.Calls)
		}
	})

	t.Run("should sleep after each count", func(t *testing.T) {
		countdownOps := &CountdownOperations{}
		expected := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		Countdown(countdownOps, countdownOps)

		if !reflect.DeepEqual(countdownOps.Ops, expected) {
			t.Errorf("expected %v, but got %v", expected, countdownOps.Ops)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepDuration := 5 * time.Second
	spyTime := &SpyTime{}

	configurableSleeper := &ConfigurableSleeper{sleepDuration, spyTime.Sleep}
	configurableSleeper.Sleep()

	if spyTime.durationSlept != sleepDuration {
		t.Errorf("should have slept for %v, but only slept %v", sleepDuration, spyTime.durationSlept)
	}
}
