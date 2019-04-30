package internal

import (
    "testing"
    "net/http"
    "io/ioutil"
    "bytes"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "time"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req), nil
}

func NewTestClient(fn RoundTripFunc, respTime time.Duration) *http.Client {
    return &http.Client{
        Transport: RoundTripFunc(fn),
        Timeout:   respTime}
}

func client(timeout time.Duration, resp *http.Response) *http.Client {
    return NewTestClient(
        func(req *http.Request) *http.Response {
            if resp != nil {
                return resp
            }

            return &http.Response{
                StatusCode: http.StatusOK,
                Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
                Header:     make(http.Header)}
        }, timeout)
}

func TestDoRequest(t *testing.T) {
    cases := map[string]struct {
        req    *internal.Request
        client *http.Client
    }{
        "valid request": {
            req:    &internal.Request{},
            client: client(30*time.Second, nil),
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            api := internal.API{Req: c.req, Client: c.client}
            res, err := api.DoRequest()

            if err != nil {
                t.Errorf("api raise error. err: %v", err)
            }

            if res.StatusCode != http.StatusOK {
                t.Errorf("api failed. response status: %d", res.StatusCode)
            }
        })
    }
}
