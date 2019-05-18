package internal

import (
    "strings"
    "net/http"
    "time"
    "bytes"
    "log"
    "io/ioutil"
)

/**
Multiple headers define like this:
Accept: application/json,text/csv,application/xml

not like this:
Accept: application/json
Accept: text/csv
Accept: application/xml
*/
type Request struct {
    Method   string
    Endpoint string
    Body     string
    Headers  map[string][]string
    Timeout  time.Duration
    Api      API
}

func BuildHeader(headers ...string) map[string][]string {
    headersMap := make(http.Header)
    for _, h := range headers {
        h := strings.Split(h, ":")
        k, v := strings.TrimSpace(h[0]), strings.TrimSpace(h[1])
        headersMap[k] = strings.Split(strings.TrimSpace(v), ",")
    }
    return headersMap
}

func getClient(timeout time.Duration) *http.Client {
    return &http.Client{Timeout: timeout * time.Second}
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

    api := NewApi(request, getClient(req.Timeout))
    res, execTime, err := api.DoRequest()
    if err != nil {
        log.Fatal(err)
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    return &Result{StatusCode: res.StatusCode, Body: b, Time: execTime}
}
