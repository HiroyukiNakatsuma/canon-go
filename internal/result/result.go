package result

import "time"

type Result struct {
    StatusCode   int
    ResponseBody []byte
    ResponseTime time.Duration
}
