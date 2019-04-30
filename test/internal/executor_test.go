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

func NewApiMock() *ApiMock {
    return &ApiMock{}
}

func (api *ApiMock) DoRequest() (*http.Response, error) {
    return &http.Response{
        StatusCode: http.StatusOK,
        Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
        Header:     make(http.Header)},
        nil
}

func NewTestExecutor(req *internal.Request) *internal.Executor {
    return &internal.Executor{Req: req, Api: NewApiMock()}
}

func TestDo(t *testing.T) {
    body := `{"greet":"Hello World!"}`
    httpMethod := "POST"
    endpoint := "http://example.com/"
    header1 := "content-type: application/json"
    header2 := "Authorization: Bearer tokenExample"
    var req = internal.Request{Method: httpMethod, Endpoint: endpoint, Body: body, Headers: internal.BuildHeader(header1, header2)}
    NewTestExecutor(&req).Do()
}
