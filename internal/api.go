package internal

import (
    "net/http"
    "bytes"
    "log"
    "time"
)

type API struct {
    Req    *Request
    Client *http.Client
}

func getDefaultClient() *http.Client {
    return &http.Client{Timeout: 30 * time.Second}
}

func (api *API) DoRequest() *http.Response {
    log.Printf("Request: %v", *api.Req)

    if api.Client == nil {
        api.Client = getDefaultClient()
    }

    req, err := http.NewRequest(api.Req.Method, api.Req.Endpoint, bytes.NewBuffer([]byte(api.Req.Body)))
    for k, v := range api.Req.Headers {
        req.Header.Add(k, v)
    }

    res, err := api.Client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    return res
}
