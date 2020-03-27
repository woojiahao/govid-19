package server

import (
  "fmt"
  "time"
)

// The data sets are updated every day at GMT +0 00:00:00

const INTERVAL_PERIOD = 24 * time.Hour

const HOUR_TO_TICK = 0
const MINUTE_TO_TICK = 0
const SECOND_TO_TICK = 0

type jobTicker struct {
  timer *time.Timer
}

func Run(f func()) {
  jobTicker := &jobTicker{}
  jobTicker.updateTimer()
  for {
    <-jobTicker.timer.C
    fmt.Println(time.Now(), " - just ticked")
    f()
    jobTicker.updateTimer()
  }
}

func (t *jobTicker) updateTimer() {
  now := time.Now()
  loc := time.FixedZone("UTC-8", 0)
  nextTick := time.Date(
    now.Year(),
    now.Month(),
    now.Day(),
    HOUR_TO_TICK,
    MINUTE_TO_TICK,
    SECOND_TO_TICK,
    0,
    loc,
  )
  if !nextTick.After(now) {
    nextTick = nextTick.Add(INTERVAL_PERIOD)
  }
  fmt.Println(nextTick, " - next tick")
  diff := nextTick.Sub(now)
  if t.timer == nil {
    t.timer = time.NewTimer(diff)
  } else {
    t.timer.Reset(diff)
  }
}
