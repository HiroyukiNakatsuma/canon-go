package data_output

import (
    "testing"
    "os"
    "fmt"
    "net/http"
    "time"

    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_output"
    "github.com/HiroyukiNakatsuma/canon-go/internal/result"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

func TestNewJsonOutput(t *testing.T) {
    cases := map[string]struct {
        outputFilepath     string
        expectHasError     bool
        expectErrorMessage string
    }{
        "valid filepath": {
            outputFilepath:     "./report.json",
            expectHasError:     false,
            expectErrorMessage: "",
        },
        "invalid filepath": {
            outputFilepath:     "./hoge/fuga.json",
            expectHasError:     true,
            expectErrorMessage: "invalid output filepath",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            _, err := data_output.NewJsonOutput(c.outputFilepath)

            if c.expectHasError && err == nil {
                t.Errorf("fail validation")
            }
            if err != nil && c.expectErrorMessage != err.Error() {
                t.Errorf("invalid error message")
            }
            if err == nil {
                CleanResultFile(c.outputFilepath)
            }
        })
    }
}

func TestOutputReport(t *testing.T) {
    cases := map[string]struct {
        actions        []action.Action
        outputFilepath string
        expectHasError bool
    }{
        "single request": {
            actions: []action.Action{
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
            },
            outputFilepath: "./output.json",
            expectHasError: false,
        },
        "single sleep": {
            actions:        []action.Action{action.NewSleep(5)},
            outputFilepath: "./output.json",
            expectHasError: false,
        },
        "request with sleep": {
            actions: []action.Action{
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
            },
            outputFilepath: "./output.json",
            expectHasError: false,
        },
        "invalid output filepath": {
            actions: []action.Action{
                action.NewRequest(
                    http.MethodGet,
                    `http://localhost:80?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
            },
            outputFilepath: "./invalid_filepath.json",
            expectHasError: true,
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            output, _ := data_output.NewJsonOutput("./output.json")
            output.OutputReport(c.actions)

            if !c.expectHasError && !existsReportFile(c.outputFilepath) {
                t.Errorf("not exists report file.")
            }

            CleanResultFile("./output.json")
        })
    }
}

func TestSummarizeByAction(t *testing.T) {
    cases := map[string]struct {
        actions                   []action.Action
        expectLabel               []string
        expectResponseTimeAverage []string
        expectErrorRate           []string
    }{
        "simple requests": {
            actions: []action.Action{
                &action.Request{
                    Method:  http.MethodGet,
                    Url:     "http://localhost:80",
                    Body:    "",
                    Headers: map[string]string{},
                    Results: []*result.Result{
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 700,
                        },
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 300,
                        },
                        {
                            StatusCode:   http.StatusNoContent,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 400,
                        },
                        {
                            StatusCode:   http.StatusBadRequest,
                            ResponseBody: []byte("NG!!!"),
                            ResponseTime: time.Millisecond * 600,
                        },
                    },
                },
                &action.Request{
                    Method:  http.MethodPost,
                    Url:     "http://localhost:80",
                    Body:    `{"hoge":"fuga""}`,
                    Headers: map[string]string{},
                    Results: []*result.Result{
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 800,
                        },
                        {
                            StatusCode:   http.StatusInternalServerError,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 600,
                        },
                    },
                },
            },
            expectLabel: []string{
                "GET http://localhost:80",
                "POST http://localhost:80",
            },
            expectResponseTimeAverage: []string{
                "0.50s",
                "0.70s",
            },
            expectErrorRate: []string{
                "25%",
                "50%",
            },
        },
        "requests with sleep": {
            actions: []action.Action{
                &action.Request{
                    Method:  http.MethodGet,
                    Url:     "http://localhost:80",
                    Body:    "",
                    Headers: map[string]string{},
                    Results: []*result.Result{
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 700,
                        },
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 300,
                        },
                        {
                            StatusCode:   http.StatusNoContent,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 400,
                        },
                        {
                            StatusCode:   http.StatusBadRequest,
                            ResponseBody: []byte("NG!!!"),
                            ResponseTime: time.Millisecond * 600,
                        },
                    },
                },
                &action.Sleep{Duration: 10},
                &action.Request{
                    Method:  http.MethodPost,
                    Url:     "http://localhost:80",
                    Body:    `{"hoge":"fuga""}`,
                    Headers: map[string]string{},
                    Results: []*result.Result{
                        {
                            StatusCode:   http.StatusOK,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 800,
                        },
                        {
                            StatusCode:   http.StatusInternalServerError,
                            ResponseBody: []byte("OK!"),
                            ResponseTime: time.Millisecond * 600,
                        },
                    },
                },
            },
            expectLabel: []string{
                "GET http://localhost:80",
                "POST http://localhost:80",
            },
            expectResponseTimeAverage: []string{
                "0.50s",
                "0.70s",
            },
            expectErrorRate: []string{
                "25%",
                "50%",
            },
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            output, _ := data_output.NewJsonOutput("./output.json")
            for i, summary := range output.SummarizeByAction(c.actions) {
                if summary.Label != c.expectLabel[i] {
                    t.Errorf("invalid summary Label.")
                }

                if summary.ResponseTimeAverage != c.expectResponseTimeAverage[i] {
                    t.Errorf("invalid summary ResponseTimeAverage.")
                }

                if summary.ErrorRate != c.expectErrorRate[i] {
                    t.Errorf("invalid summary ErrorRate.")
                }
            }
            if existsReportFile("./output.json") {
                CleanResultFile("./output.json")
            }
        })
    }
}

func existsReportFile(filepath string) bool {
    _, err := os.Stat(filepath)
    return err == nil
}

func CleanResultFile(filepath string) {
    err := os.Remove(filepath)
    if err != nil {
        fmt.Println(err)
    }
}
