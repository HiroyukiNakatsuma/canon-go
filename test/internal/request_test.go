package internal

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

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
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid POST request": {
            Request: internal.NewRequest(
                http.MethodPost,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid PUT request": {
            Request: internal.NewRequest(
                http.MethodPut,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid DELETE request": {
            Request: internal.NewRequest(
                http.MethodDelete,
                `http://example.com`,
                `{"greet":"Hello World!"}`,
                internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                mock.NewMockClient(30, nil)),
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
