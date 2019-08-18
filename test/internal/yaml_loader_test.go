package internal

import (
    "testing"
    "io/ioutil"
    "log"
    "os"
    "time"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

func TestLoadActions(t *testing.T) {
    cases := map[string]struct {
        inputContent         []byte
        inputFilepath        string
        expectActions        []internal.Action
        expectHasError       bool
        expectedErrorMessage string
    }{
        "request only required param": {
            inputContent: []byte(`
actions:
  - request:
      method: GET
      url: http://example.com/
`),
            inputFilepath: "./input.yml",
            expectActions: []internal.Action{
                &internal.Request{
                    Method: "GET",
                    Url:    "http://example.com/",
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
            expectActions: []internal.Action{
                &internal.Sleep{Duration: 10 * time.Second},
                &internal.Sleep{Duration: 20 * time.Second},
            },
            expectHasError:       false,
            expectedErrorMessage: "",
        },
        "request has body": {
            inputContent: []byte(`
actions:
  - request:
      method: POST
      url: http://example.com/
      body: '{"hoge":1,"fuga":2}'
`),
            inputFilepath: "./input.yml",
            expectActions: []internal.Action{
                &internal.Request{
                    Method: "POST",
                    Url:    "http://example.com/",
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
      url: http://example.com/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Content-Type: application/json
        Authorization: Bearer token
`),
            inputFilepath: "./input.yml",
            expectActions: []internal.Action{
                &internal.Request{
                    Method: "POST",
                    Url:    "http://example.com/",
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
      url: http://example.com/
  - sleep: 10
  - request:
      method: POST
      url: http://example.com/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Content-Type: application/json
        Authorization: Bearer token
`),
            inputFilepath: "./input.yml",
            expectActions: []internal.Action{
                &internal.Request{
                    Method: "GET",
                    Url:    "http://example.com/",
                },
                &internal.Sleep{Duration: 10 * time.Second},
                &internal.Request{
                    Method: "POST",
                    Url:    "http://example.com/",
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
      url: http://example.com/
      body: '{"hoge":1,"fuga":2}'
      headers:
        Accept: application/json
        Accept: text/html
`),
            inputFilepath: "./input.yml",
            expectActions: []internal.Action{
                &internal.Request{
                    Method:  "POST",
                    Url:     "http://example.com/",
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
      url: http://example.com/
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
      url: http://example.com/
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

            actions, err := internal.NewYamlLoader(c.inputFilepath).LoadActions()

            if err != nil && !c.expectHasError {
                t.Errorf("must raise error")
            }
            if err != nil && err.Error() != c.expectedErrorMessage {
                t.Errorf("invalid error message")
            }

            for i, action := range actions {
                request, isRequest := action.(*internal.Request)
                if isRequest {
                    expectRequest := c.expectActions[i].(*internal.Request)
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
                    sleep := action.(*internal.Sleep).Duration.Seconds()
                    expectSleep := c.expectActions[i].(*internal.Sleep).Duration.Seconds()
                    if sleep != expectSleep {
                        t.Errorf("invalid sleep duration")
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
