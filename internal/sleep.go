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
    return &Sleep{Duration: duration}
}

func (slp *Sleep) Do() {
    log.Printf("sleep started.")
    time.Sleep(slp.Duration * time.Second)
    log.Printf("sleep finished.")
}

func (slp *Sleep) GetResults() []*Result {
    return slp.Results
}
