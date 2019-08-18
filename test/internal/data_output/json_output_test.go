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

func TestOutputReport(t *testing.T) {
    cases := map[string]struct {
        actions              []action.Action
        expectHasError       bool
        expectedErrorMessage string
    }{
        "single request": {
            actions: []action.Action{
                action.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "single sleep": {
            actions:              []action.Action{action.NewSleep(5)},
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request with sleep": {
            actions: []action.Action{
                action.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
                action.NewSleep(5),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            data_output.NewJsonOutput().OutputReport(c.actions)

            if !existsReportFile() {
                t.Errorf("not exists report file.")
            }

            CleanResultFile()
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
                    Url:     "http://example.com",
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
                    Url:     "http://example.com",
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
                "GET http://example.com",
                "POST http://example.com",
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
                    Url:     "http://example.com",
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
                    Url:     "http://example.com",
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
                "GET http://example.com",
                "POST http://example.com",
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
            summaries := data_output.NewJsonOutput().SummarizeByAction(c.actions)
            for i, summary := range summaries {
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
        })
    }
}

func existsReportFile() bool {
    _, err := os.Stat("./report.json")
    return err == nil
}

func CleanResultFile() {
    err := os.Remove("./report.json")
    if err != nil {
        fmt.Println(err)
    }
}
