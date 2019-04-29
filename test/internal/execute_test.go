package internal

import (
    "testing"
    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "net/http"
    "io/ioutil"
    "bytes"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
    return &http.Client{
        Transport: RoundTripFunc(fn),
    }
}

func TestExecute(t *testing.T) {
    body := `{"greet":"Hello World!"}`
    httpMethod := "POST"
    endpoint := "http://example.com/"
    header1 := "content-type: application/json"
    header2 := "Authorization: Bearer tokenExample"
    var req = internal.Request{Method: httpMethod, Endpoint: endpoint, Body: body, Headers: internal.BuildHeader(header1, header2)}

    client := NewTestClient(func(req *http.Request) *http.Response {
        return &http.Response{
            StatusCode: 200,
            Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
            Header:     make(http.Header)}
    })

    api := internal.API{Req: &req, Client: client}
    res := api.Execute()
    if res != "OK" {
        t.Errorf("api failed. res: %s", res)
    }
}
