package internal

import (
    "net/http"
    "bytes"
    "log"
    "time"
)

type API interface {
    DoRequest() (*http.Response, error)
}

type Api struct {
    Req    *Request
    Client *http.Client
}

func NewApi(req *Request, client *http.Client) *Api {
    return &Api{Req: req, Client: client}
}

func getDefaultClient() *http.Client {
    return &http.Client{Timeout: 30 * time.Second}
}

func (api *Api) DoRequest() (*http.Response, error) {
    log.Printf("Request: %v", *api.Req)

    if api.Client == nil {
        api.Client = getDefaultClient()
    }

    req, err := http.NewRequest(api.Req.Method, api.Req.Endpoint, bytes.NewBuffer([]byte(api.Req.Body)))
    for k, vs := range api.Req.Headers {
        for _, v := range vs {
            req.Header.Add(k, v)
        }
    }

    res, err := api.Client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    return res, err
}
