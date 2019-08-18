package internal

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

func TestExecute(t *testing.T) {
    cases := map[string]struct {
        dataInput            internal.DataInput
        dataOutput           internal.DataOutput
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid requests": {
            dataInput: mock.NewYamlLoaderMock(
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPost,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    nil,
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPut,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    nil,
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodDelete,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    nil,
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput:           mock.NewJsonOutputMock(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with headers": {
            dataInput: mock.NewYamlLoaderMock(
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPost,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPut,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodDelete,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput:           mock.NewJsonOutputMock(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with sleep": {
            dataInput: mock.NewYamlLoaderMock(
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodPost,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodPut,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodDelete,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput:           mock.NewJsonOutputMock(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            executor := internal.NewExecutor(c.dataInput, c.dataOutput)
            executor.Execute()
        })
    }
}
