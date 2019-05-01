package internal

import (
    "log"
    "io/ioutil"
)

type Executor struct {
    Apis []API
}

func NewExecutor(reqs ...*Request) *Executor {
    var apis []API
    for _, req := range reqs {
        apis = append(apis, NewApi(req, nil))
    }
    return &Executor{Apis: apis}
}

func (e *Executor) Do() {
    for _, api := range e.Apis {
        log.Printf("req: %v", api)

        res, time, err := api.DoRequest()
        if err != nil {
            log.Fatal(err)
        }

        b, err := ioutil.ReadAll(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        var result = Result{StatusCode: res.StatusCode, Body: b, Time: time}
        log.Printf("Response Status: %d", result.StatusCode)
        log.Printf("Response Body: %s", result.Body)
        log.Printf("Response Time: %fs", result.Time.Seconds())
    }
}
