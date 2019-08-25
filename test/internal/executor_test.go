package internal

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/internal/config"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_input"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_output"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

func TestExecute(t *testing.T) {
    cases := map[string]struct {
        dataInput            data_input.DataInput
        dataOutput           data_output.DataOutput
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid requests": {
            dataInput: mock.NewYamlLoaderMock(
                &config.Config{Threads: 1, Loop: 1},
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodPost,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    nil,
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodPut,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    nil,
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodDelete,
                    `http://localhost:80`,
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
                &config.Config{Threads: 1, Loop: 1},
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodPost,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodPut,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewRequest(
                    http.MethodDelete,
                    `http://localhost:80`,
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
                &config.Config{Threads: 1, Loop: 1},
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodPost,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodPut,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodDelete,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput:           mock.NewJsonOutputMock(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "loop 3 requests & sleep": {
            dataInput: mock.NewYamlLoaderMock(
                &config.Config{Threads: 1, Loop: 3},
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodPost,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodPut,
                    `http://localhost:80`,
                    `{"greet":"Hello World!"}`,
                    map[string]string{"content-type": "application/json", "Authorization": "Bearer tokenExample"},
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
                action.NewRequest(
                    http.MethodDelete,
                    `http://localhost:80`,
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
