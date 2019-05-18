package internal

import (
    "testing"
    "net/http"
    "io/ioutil"
    "bytes"
    "time"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
    return f(req), nil
}

func NewTestClient(fn RoundTripFunc, respTime time.Duration) *http.Client {
    return &http.Client{
        Transport: RoundTripFunc(fn),
        Timeout:   respTime,
    }
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

func NewTestRequest(method, endpoint, body string, headers ...string) *http.Request {
    request, _ := http.NewRequest(method, endpoint, bytes.NewBuffer([]byte(body)))
    hs := internal.BuildHeader(headers...)
    for k, vs := range hs {
        for _, v := range vs {
            request.Header.Add(k, v)
        }
    }
    return request
}

func TestDoRequest(t *testing.T) {
    cases := map[string]struct {
        req                  *http.Request
        client               *http.Client
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid GET request": {
            req: NewTestRequest(
                http.MethodGet,
                `http://example.com?greet="Hello World!"`,
                ``,
                `content-type: application/json`,
                `Authorization: Bearer tokenExample`),
            client:               client(30*time.Second, nil),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid POST request": {
            req: NewTestRequest(
                http.MethodPost,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                `content-type: application/json`,
                `Authorization: Bearer tokenExample`),
            client:               client(30*time.Second, nil),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid PUT request": {
            req: NewTestRequest(
                http.MethodPut,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                `content-type: application/json`,
                `Authorization: Bearer tokenExample`),
            client:               client(30*time.Second, nil),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid DELETE request": {
            req: NewTestRequest(
                http.MethodDelete,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                `content-type: application/json`,
                `Authorization: Bearer tokenExample`),
            client:               client(30*time.Second, nil),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            api := internal.NewApi(c.req, c.client)
            res, _, err := api.DoRequest()

            if c.expectHasError {
                if err == nil {
                    t.Errorf("api not raise error.")
                }

                if err.Error() != c.expectedErrorMessage {
                    t.Errorf("unexpected error message. expected: %s, actual: %s", c.expectedErrorMessage, err.Error())
                }
            }

            if res.StatusCode != http.StatusOK {
                t.Errorf("api failed. response status: %d", res.StatusCode)
            }
        })
    }
}
