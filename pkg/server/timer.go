package server

import (
  "fmt"
  "time"
)

// The data sets are updated every day at GMT +0 00:00:00
var location = time.FixedZone("GMT+0", 0)

const (
  interval = 24 * time.Hour
  hour     = 0
  minute   = 0
  second   = 0
)

type jobTicker struct {
  timer *time.Timer
}

// Updates the timer to include a new tick once the original is up.
func (t *jobTicker) updateTimer() {
  now := time.Now().In(location)

  nextTick := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, location)

  if !nextTick.After(now) {
    nextTick = nextTick.Add(interval)
  }

  diff := nextTick.Sub(now)

  if t.timer == nil {
    t.timer = time.NewTimer(diff)
  } else {
    t.timer.Reset(diff)
  }
}

func Run(f func()) {
  jobTicker := &jobTicker{}
  jobTicker.updateTimer()

  for {
    <-jobTicker.timer.C
    fmt.Println(time.Now().In(location), " - just ticked")
    f()
    jobTicker.updateTimer()
  }
}
