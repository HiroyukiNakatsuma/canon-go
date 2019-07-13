package mock

import (
    "net/http"
    "time"
    "io/ioutil"
    "bytes"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req), nil
}

func client(fn RoundTripFunc, respTime time.Duration) *http.Client {
    return &http.Client{
        Transport: RoundTripFunc(fn),
        Timeout:   respTime,
    }
}

func NewMockClient(timeout time.Duration, resp *http.Response) *http.Client {
    return client(
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
