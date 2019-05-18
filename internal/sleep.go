package internal

import "time"

type Sleep struct {
    Duration time.Duration
}

func (slp *Sleep) Do() *Result {
    time.Sleep(slp.Duration * time.Second)
    return nil
}
