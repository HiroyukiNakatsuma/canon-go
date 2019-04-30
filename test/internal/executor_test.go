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
    cases := map[string]struct {
        req                  *internal.Request
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid GET request": {
            req: &internal.Request{
                Method:   http.MethodGet,
                Endpoint: `http://example.com?greet="Hello World!"`,
                Body:     ``,
                Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid POST request": {
            req: &internal.Request{
                Method:   http.MethodPost,
                Endpoint: `http://example.com`,
                Body:     `{"greet":"Hello World!"}`,
                Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid PUT request": {
            req: &internal.Request{
                Method:   http.MethodPut,
                Endpoint: `http://example.com`,
                Body:     `{"greet":"Hello World!"}`,
                Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid DELETE request": {
            req: &internal.Request{
                Method:   http.MethodDelete,
                Endpoint: `http://example.com`,
                Body:     `{"greet":"Hello World!"}`,
                Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            NewTestExecutor(c.req).Do()
        })
    }
}
