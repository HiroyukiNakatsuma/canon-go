package action

import (
    "net/http"
    "time"
    "bytes"
    "log"
    "io/ioutil"

    "github.com/HiroyukiNakatsuma/canon-go/internal/result"
)

type Request struct {
    Method  string
    Url     string
    Body    string
    Headers map[string]string
    Client  *http.Client
    Results []*result.Result
}

func NewRequest(method string, url string, body string, headers map[string]string, client *http.Client) *Request {
    return &Request{
        Method:  method,
        Url:     url,
        Body:    body,
        Headers: headers,
        Client:  client,
    }
}

func (req *Request) Do() {
    request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer([]byte(req.Body)))
    if err != nil {
        log.Fatal(err)
    }
    for k, v := range req.Headers {
        request.Header.Add(k, v)
    }

    log.Printf("start request.")

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

    rslt := &result.Result{StatusCode: res.StatusCode, ResponseBody: b, ResponseTime: execTime}

    log.Printf("Response Status: %d", rslt.StatusCode)
    log.Printf("Response Time: %fs", rslt.ResponseTime.Seconds())

    req.Results = append(req.Results, rslt)
}

func (req *Request) GetResults() []*result.Result {
    return req.Results
}
