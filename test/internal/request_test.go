package internal

import (
    "testing"
    "net/http"
    "time"
    "io/ioutil"
    "bytes"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
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

func TestDo(t *testing.T) {
    cases := map[string]struct {
        Request              *internal.Request
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid GET Request": {
            Request: internal.NewRequest(
                http.MethodGet,
                `http://example.com?greet="Hello World!"`,
                ``,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid POST request": {
            Request: internal.NewRequest(
                http.MethodPost,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid PUT request": {
            Request: internal.NewRequest(
                http.MethodPut,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid DELETE request": {
            Request: internal.NewRequest(
                http.MethodDelete,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            result := c.Request.Do()
            if result.StatusCode != http.StatusOK {
                t.Errorf("api failed. response status: %d", result.StatusCode)
            }
        })
    }
}
