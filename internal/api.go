package internal

import (
    "net/http"
    "log"
    "time"
)

type API interface {
    DoRequest() (*http.Response, time.Duration, error)
}

type Api struct {
    Req    *http.Request
    Client *http.Client
}

func NewApi(req *http.Request, client *http.Client) *Api {
    return &Api{Req: req, Client: client}
}

func (api *Api) DoRequest() (*http.Response, time.Duration, error) {
    log.Printf("Request: %v", *api.Req)

    start := time.Now()

    res, err := api.Client.Do(api.Req)
    defer res.Body.Close()

    duration := time.Now().Sub(start)

    if err != nil {
        log.Fatal(err)
    }

    return res, duration, err
}
