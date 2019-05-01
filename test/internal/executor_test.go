package internal

import (
    "testing"
    "net/http"
    "bytes"
    "io/ioutil"
    "time"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

type ApiMock struct {
    Req    *internal.Request
    Client *http.Client
}

func NewApiMock() *ApiMock {
    return &ApiMock{}
}

func (api *ApiMock) DoRequest() (*http.Response, time.Duration, error) {
    return &http.Response{
        StatusCode: http.StatusOK,
        Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
        Header:     make(http.Header)},
        100 * time.Millisecond,
        nil
}

func NewTestExecutor(reqs []*internal.Request) *internal.Executor {
    var apis []internal.API
    for range reqs {
        apis = append(apis, NewApiMock())
    }
    return &internal.Executor{Apis: apis}
}

func TestDo(t *testing.T) {
    cases := map[string]struct {
        reqs                 []*internal.Request
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid requests": {
            reqs: []*internal.Request{
                {
                    Method:   http.MethodGet,
                    Endpoint: `http://example.com?greet="Hello World!"`,
                    Body:     ``,
                },
                {
                    Method:   http.MethodPost,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                },
                {
                    Method:   http.MethodPut,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                },
                {
                    Method:   http.MethodDelete,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with headers": {
            reqs: []*internal.Request{
                {
                    Method:   http.MethodGet,
                    Endpoint: `http://example.com?greet="Hello World!"`,
                    Body:     ``,
                    Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                },
                {
                    Method:   http.MethodPost,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                    Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                },
                {
                    Method:   http.MethodPut,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                    Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                },
                {
                    Method:   http.MethodDelete,
                    Endpoint: `http://example.com`,
                    Body:     `{"greet":"Hello World!"}`,
                    Headers:  internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            NewTestExecutor(c.reqs).Do()
        })
    }
}
