package internal

import "time"

type Sleep struct {
    Duration time.Duration
}

func NewSleep(duration time.Duration) *Sleep {
    return &Sleep{Duration: duration}
}

func (slp *Sleep) Do() *Result {
    time.Sleep(slp.Duration * time.Second)
    return nil
}
