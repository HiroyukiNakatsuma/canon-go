package internal

import (
    "time"
    "log"
)

type Sleep struct {
    Duration time.Duration
    Results  []*Result
}

func NewSleep(duration time.Duration) *Sleep {
    return &Sleep{Duration: duration * time.Second}
}

func (slp *Sleep) Do() {
    log.Printf("sleep started.")
    time.Sleep(slp.Duration)
    log.Printf("sleep finished.")
}

func (slp *Sleep) GetResults() []*Result {
    return slp.Results
}
