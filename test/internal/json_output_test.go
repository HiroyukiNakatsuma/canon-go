package internal

import (
    "testing"
    "os"
    "fmt"
    "net/http"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/test/mock"
)

func TestOutputReport(t *testing.T) {
    cases := map[string]struct {
        actions              []internal.Action
        expectHasError       bool
        expectedErrorMessage string
    }{
        "single request": {
            actions: []internal.Action{
                internal.NewRequest(
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
            actions:              []internal.Action{internal.NewSleep(5)},
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request with sleep": {
            actions: []internal.Action{
                internal.NewRequest(
                    http.MethodGet,
                    `http://example.com?greet="Hello World!"`,
                    ``,
                    nil,
                    mock.NewMockClient(30, nil)),
                internal.NewSleep(5),
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            internal.NewJsonOutput().OutputReport(c.actions)

            if !existsReportFile() {
                t.Errorf("not exists report file.")
            }

            afterTest()
        })
    }
}

func existsReportFile() bool {
    _, err := os.Stat("./report.json")
    return err == nil
}

func afterTest() {
    err := os.Remove("./report.json")
    if err != nil {
        fmt.Println(err)
    }
}
