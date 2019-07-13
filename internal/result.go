package internal

import "time"

type Result struct {
    Request    *Request
    StatusCode int
    Body       []byte
    Time       time.Duration
}
