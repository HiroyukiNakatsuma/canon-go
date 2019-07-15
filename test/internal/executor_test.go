package internal

import (
    "testing"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

type YamlLoaderMock struct {
    actions []internal.Action
}

func NewYamlLoaderMock(actions ...internal.Action) *YamlLoaderMock {
    return &YamlLoaderMock{actions: actions}
}

func (yamlLoader *YamlLoaderMock) LoadConfig() *internal.ActionConfig {
    return &internal.ActionConfig{Threads: 1, Loop: 1}
}

func (yamlLoader *YamlLoaderMock) LoadActions() []internal.Action {
    return yamlLoader.actions
}

func TestExecute(t *testing.T) {
    cases := map[string]struct {
        dataInput            internal.DataInput
        dataOutput           internal.DataOutput
        expectHasError       bool
        expectedErrorMessage string
    }{
        "valid requests": {
            dataInput: NewYamlLoaderMock(
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
            dataOutput: internal.NewJsonOutput(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with headers": {
            dataInput: NewYamlLoaderMock(
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPost,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodPut,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewRequest(
                    http.MethodDelete,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput: internal.NewJsonOutput(),
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "valid requests with sleep": {
            dataInput: NewYamlLoaderMock(
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodPost,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodPut,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
                internal.NewRequest(
                    http.MethodDelete,
                    `http://example.com`,
                    `{"greet":"Hello World!"}`,
                    internal.BuildHeader(`content-type: application/json`, `Authorization: Bearer tokenExample`),
                    mock.NewMockClient(30, nil)),
            ),
            dataOutput: internal.NewJsonOutput(),
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
