package internal

import (
    "testing"
    "net/http"
    "bytes"
    "io/ioutil"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

type ApiMock struct {
    Req    *internal.Request
    Client *http.Client
}

func NewApiMock(req *internal.Request, client *http.Client) *ApiMock {
    return &ApiMock{Req: req, Client: client}
}

func (api *ApiMock) DoRequest() (*http.Response, error) {
    return &http.Response{
        StatusCode: http.StatusOK,
        Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
        Header:     make(http.Header)},
        nil
}

func TestExecute(t *testing.T) {
}
