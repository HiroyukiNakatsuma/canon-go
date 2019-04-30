package internal

import (
    "log"
    "io/ioutil"
)

type Executor struct {
    Req *Request
    Api API
}

func NewExecutor(req *Request) *Executor {
    return &Executor{Req: req, Api: NewApi(req, nil)}
}

func (e *Executor) Do() {
    log.Printf("req: %v", e.Req)

    res, time, err := e.Api.DoRequest()
    if err != nil {
        log.Fatal(err)
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Response Status: %d", res.StatusCode)
    log.Printf("Response Body: %s", b)
    log.Printf("Response Time: %fs", time.Seconds())
}
