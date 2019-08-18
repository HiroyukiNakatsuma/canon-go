package action

import (
    "time"
    "log"

    "github.com/HiroyukiNakatsuma/canon-go/internal/result"
)

type Sleep struct {
    Duration time.Duration
    Results  []*result.Result
}

func NewSleep(duration time.Duration) *Sleep {
    return &Sleep{Duration: duration * time.Second}
}

func (slp *Sleep) Do() {
    log.Printf("sleep started.")
    time.Sleep(slp.Duration)
    log.Printf("sleep finished.")
}

func (slp *Sleep) GetResults() []*result.Result {
    return slp.Results
}
