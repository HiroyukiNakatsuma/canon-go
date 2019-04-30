package internal

import "time"

type Result struct {
    StatusCode int
    Body       []byte
    Time       time.Duration
}
