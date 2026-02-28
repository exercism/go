package splitsecondstopwatch

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func formatTime(time int) string {
	return fmt.Sprintf("%02d:%02d:%02d", time/60/60, time/60%60, time%60)
}

func parseTime(time string) int {
	parts := strings.Split(time, ":")
	sum := 0
	for _, part := range parts {
		v, _ := strconv.Atoi(part)
		sum *= 60
		sum += v
	}
	return sum
}

type SplitSecondStopwatch struct {
	state    string
	laps     []int
	elapsed  int
	lapStart int
	started  bool
}

func (sss *SplitSecondStopwatch) Start() error {
	if sss.state == "running" {
		return errors.New("cannot start an already running stopwatch")
	}
	sss.started = true
	sss.state = "running"
	return nil
}

func (sss *SplitSecondStopwatch) Stop() error {
	if sss.state != "running" {
		return errors.New("cannot stop a stopwatch that is not running")
	}
	sss.state = "stopped"
	return nil
}

func (sss *SplitSecondStopwatch) Reset() error {
	if sss.state != "stopped" {
		return errors.New("cannot reset a stopwatch that is not stopped")
	}
	sss.laps = []int{}
	sss.state = "ready"
	sss.elapsed = 0
	sss.lapStart = 0
	sss.started = false
	return nil
}

func (sss *SplitSecondStopwatch) Lap() error {
	if sss.state != "running" {
		return errors.New("cannot lap a stopwatch that is not running")
	}
	sss.state = "running"
	sss.laps = append(sss.laps, sss.elapsed-sss.lapStart)
	sss.lapStart = sss.elapsed
	return nil
}

func (sss *SplitSecondStopwatch) AdvanceTime(by string) {
	if sss.state == "running" {
		sss.elapsed += parseTime(by)
	}
}

func (sss *SplitSecondStopwatch) State() string {
	return sss.state
}

func (sss *SplitSecondStopwatch) CurrentLap() string {
	return formatTime(sss.elapsed - sss.lapStart)
}

func (sss *SplitSecondStopwatch) Total() string {
	return formatTime(sss.elapsed)
}

func (sss *SplitSecondStopwatch) PreviousLaps() []string {
	out := []string{}
	for _, time := range sss.laps {
		out = append(out, formatTime(time))
	}
	return out
}

func NewSplitSecondStopwatch() *SplitSecondStopwatch {
	return &SplitSecondStopwatch{
		state: "ready",
		laps:  []int{},
	}
}
