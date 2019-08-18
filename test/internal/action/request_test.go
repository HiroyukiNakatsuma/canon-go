package action

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

func TestDo(t *testing.T) {
    cases := map[string]struct {
        Request              *action.Request
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid GET Request": {
            Request: action.NewRequest(
                http.MethodGet,
                `http://localhost:80?greet="Hello World!"`,
                ``,
                map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid POST request": {
            Request: action.NewRequest(
                http.MethodPost,
                `http://localhost:80`,
                `{"greet":"Hello World!"}`,
                map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid PUT request": {
            Request: action.NewRequest(
                http.MethodPut,
                `http://localhost:80`,
                `{"greet":"Hello World!"}`,
                map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid DELETE request": {
            Request: action.NewRequest(
                http.MethodDelete,
                `http://localhost:80`,
                `{"greet":"Hello World!"}`,
                map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                mock.NewMockClient(30, nil)),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            c.Request.Do()
        })
    }
}
