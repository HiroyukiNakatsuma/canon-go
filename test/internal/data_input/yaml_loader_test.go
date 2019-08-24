package data_input

import (
    "testing"
    "io/ioutil"
    "log"
    "os"
    "time"

    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_input"
)

func TestLoadActions(t *testing.T) {
    cases := map[string]struct {
        inputContent         []byte
        inputFilepath        string
        expectActions        []action.Action
        expectHasError       bool
        expectedErrorMessage string
    }{
        "request only required param": {
            inputContent: []byte(`
actions:
  - request:
      method: GET
      url: http://localhost:80/
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Request{
                    Method: "GET",
                    Url:    "http://localhost:80/",
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "only sleep": {
            inputContent: []byte(`
actions:
  - sleep: 10
  - sleep: 20
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Sleep{Duration: 10 * time.Second},
                &action.Sleep{Duration: 20 * time.Second},
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request has body": {
            inputContent: []byte(`
actions:
  - request:
      method: POST
      url: http://localhost:80/
      body: '{"hoge":1,"fuga":2}'
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Request{
                    Method: "POST",
                    Url:    "http://localhost:80/",
                    Body:   `{"hoge":1,"fuga":2}`,
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request has headers": {
            inputContent: []byte(`
actions:
  - request:
      method: POST
      url: http://localhost:80/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Content-Type: application/json
        Authorization: Bearer token
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Request{
                    Method: "POST",
                    Url:    "http://localhost:80/",
                    Body:   `{"hoge":1,"fuga":2}`,
                    Headers: map[string]string{
                        "Content-Type":  "application/json",
                        "Authorization": "Bearer token",
                    },
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request & sleep": {
            inputContent: []byte(`
actions:
  - request:
      method: GET
      url: http://localhost:80/
  - sleep: 10
  - request:
      method: POST
      url: http://localhost:80/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Content-Type: application/json
        Authorization: Bearer token
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Request{
                    Method: "GET",
                    Url:    "http://localhost:80/",
                },
                &action.Sleep{Duration: 10 * time.Second},
                &action.Request{
                    Method: "POST",
                    Url:    "http://localhost:80/",
                    Body:   `{"hoge":1,"fuga":2}`,
                    Headers: map[string]string{
                        "Content-Type":  "application/json",
                        "Authorization": "Bearer token",
                    },
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request has duplicated headers": {
            inputContent: []byte(`
actions:
  - request:
      method: POST
      url: http://localhost:80/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Accept: application/json
        Accept: text/html
`),
            inputFilepath: "./input.yml",
            expectActions: []action.Action{
                &action.Request{
                    Method:  "POST",
                    Url:     "http://localhost:80/",
                    Body:    `{"hoge":1,"fuga":2}`,
                    Headers: map[string]string{"Accept": "text/html"},
                },
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "invalid input filepath": {
            inputContent: []byte(`
actions:
  - request:
      method: GET
      url: http://localhost:80/
`),
            inputFilepath:        "./invalid_filepath.yml",
            expectActions:        nil,
            expectHasError:       true,
            expectedErrorMessage: "invalid filepath",
        },
        "no request or sleep": {
            inputContent: []byte(`
actions:
  - hoge: hoge
  - fuga: fuga
`),
            inputFilepath:        "./input.yml",
            expectActions:        nil,
            expectHasError:       true,
            expectedErrorMessage: "invalid input",
        },
        "request has no method": {
            inputContent: []byte(`
actions:
  - request:
      url: http://localhost:80/
`),
            inputFilepath:        "./input.yml",
            expectActions:        nil,
            expectHasError:       true,
            expectedErrorMessage: "method is required",
        },
        "request has no url": {
            inputContent: []byte(`
actions:
  - request:
      method: GET
`),
            inputFilepath:        "./input.yml",
            expectActions:        nil,
            expectHasError:       true,
            expectedErrorMessage: "url is required",
        },
    }

    for name, c := range cases {
        t.Run(name, func(t *testing.T) {
            err := ioutil.WriteFile("./input.yml", c.inputContent, 0644)
            if err != nil {
                log.Fatal(err)
            }

            input, err := data_input.NewYamlLoader(c.inputFilepath)

            if err != nil && !c.expectHasError {
                t.Errorf("must raise error")
            }
            if err != nil && err.Error() != c.expectedErrorMessage {
                t.Errorf("invalid error message")
            }

            if err == nil {
                for i, act := range input.LoadActions() {
                    request, isRequest := act.(*action.Request)
                    if isRequest {
                        expectRequest := c.expectActions[i].(*action.Request)
                        if request.Method != expectRequest.Method {
                            t.Errorf("invalid method")
                        }
                        if request.Url != expectRequest.Url {
                            t.Errorf("invalid method")
                        }
                        if request.Body != expectRequest.Body {
                            t.Errorf("invalid method")
                        }
                        for k, v := range request.Headers {
                            if v != expectRequest.Headers[k] {
                                t.Errorf("invalid header")
                            }
                        }
                    } else {
                        sleep := act.(*action.Sleep).Duration.Seconds()
                        expectSleep := c.expectActions[i].(*action.Sleep).Duration.Seconds()
                        if sleep != expectSleep {
                            t.Errorf("invalid sleep duration")
                        }
                    }
                }
            }

            err = os.Remove("./input.yml")
            if err != nil {
                log.Fatal(err)
            }
        })
    }
}
