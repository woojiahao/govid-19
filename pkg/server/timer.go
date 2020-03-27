package server

import (
  "fmt"
  "time"
)

// The data sets are updated every day at GMT +0 00:00:00

const IntervalPeriod = 24 * time.Hour

const HourToTick = 0
const MinuteToTick = 0
const SecondToTick = 0

type jobTicker struct {
  timer *time.Timer
}

// This function must be run as a goroutine to ensure that the timer does not block the main thread.
func Run(f func()) {
  jobTicker := &jobTicker{}
  jobTicker.updateTimer()
  // Infinite loop that waits for the channel to receive some results
  for {
    <-jobTicker.timer.C
    fmt.Println(time.Now(), " - just ticked")
    // Execute the given function when the interval is reached
    f()
    jobTicker.updateTimer()
  }
}

// Updates the timer to include a new tick once the original is up.
func (t *jobTicker) updateTimer() {
  now := time.Now()

  // Updates are delivered on at UTC-8 00:00:00
  loc := time.FixedZone("UTC-8", 0)

  // By default, set the next tick to the current date at 00:00:00
  nextTick := time.Date(
    now.Year(),
    now.Month(),
    now.Day(),
    HourToTick,
    MinuteToTick,
    SecondToTick,
    0,
    loc,
  )

  // If the next tick is before the current date and time, set it to be the next day (IntervalPeriod)
  // at 00:00:00
  if !nextTick.After(now) {
    nextTick = nextTick.Add(IntervalPeriod)
  }
  fmt.Println(nextTick, " - next tick")

  // To calculate when the next tick occurs, we subtract the current time from the next tick
  diff := nextTick.Sub(now)

  // Assign a timer if not yet assigned, else just reset the timer with the newly calculated difference
  if t.timer == nil {
    t.timer = time.NewTimer(diff)
  } else {
    t.timer.Reset(diff)
  }
}
