package internal

import (
    "net/http"
    "time"
    "io/ioutil"
    "bytes"
)

type ApiMock struct {
    Req    *http.Request
    Client *http.Client
}

func NewApiMock(req *http.Request, client *http.Client) *ApiMock {
    return &ApiMock{Req: req, Client: client}
}

func (api *ApiMock) DoRequest() (*http.Response, time.Duration, error) {
    return &http.Response{
        StatusCode: http.StatusOK,
        Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
        Header:     make(http.Header)},
        100 * time.Millisecond,
        nil
}
