package internal

import "time"

type Result struct {
    Request      *Request
    StatusCode   int
    ResponseBody []byte
    ResponseTime time.Duration
}
