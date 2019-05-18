package internal

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

func NewTestExecutor(actions []internal.Action) *internal.Executor {
    return &internal.Executor{Actions: actions}
}

func requests2Actions(reqs []*internal.Request) []internal.Action {
    actions := make([]internal.Action, len(reqs))
    for i, req := range reqs {
        actions[i] = req
    }
    return actions
}

func TestExecute(t *testing.T) {
    cases := map[string]struct {
        actions              []internal.Action
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid requests": {
            actions: requests2Actions([]*internal.Request{
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
            }),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with headers": {
            actions: requests2Actions([]*internal.Request{
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
            }),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            NewTestExecutor(c.actions).Execute()
        })
    }
}
