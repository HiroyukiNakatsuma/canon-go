package internal

import (
    "net/http"
    "time"
    "bytes"
    "log"
    "io/ioutil"
)

type Request struct {
    Method   string
    Endpoint string
    Body     string
    Headers  map[string][]string
    Client   *http.Client
}

func NewRequest(method string, endpoint string, body string, headers map[string][]string, client *http.Client) *Request {
    return &Request{
        Method:   method,
        Endpoint: endpoint,
        Body:     body,
        Headers:  headers,
        Client:   client,
    }
}

func (req *Request) Do() *Result {
    request, err := http.NewRequest(req.Method, req.Endpoint, bytes.NewBuffer([]byte(req.Body)))
    if err != nil {
        log.Fatal(err)
    }
    for k, vs := range req.Headers {
        for _, v := range vs {
            request.Header.Add(k, v)
        }
    }

    log.Printf("Request: %v", request)

    start := time.Now()

    res, err := req.Client.Do(request)
    defer res.Body.Close()

    execTime := time.Now().Sub(start)

    if err != nil {
        log.Fatal(err)
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    return &Result{Request: req, StatusCode: res.StatusCode, ResponseBody: b, ResponseTime: execTime}
}
