package internal

import (
    "net/http"
    "time"
    "log"
    "fmt"

    "gopkg.in/yaml.v2"
)

type yamlLoader struct{}

type input struct {
    Timeout time.Duration
    Actions []map[string]interface{}
}

func NewYamlLoader() *yamlLoader {
    return &yamlLoader{}
}

func (yamlLoader *yamlLoader) LoadConfig() *ActionConfig {
    return &ActionConfig{Threads: 1, Loop: 1}
}

func (yamlLoader *yamlLoader) LoadActions() []Action {
    yamlInput := []byte(`
timeout: 30
actions:
  - request:
      method: GET
      endpoint: http://example.com/
      body: 'greet="Hello World!"'
      headers:
          content-type: 'application/json'
          Authorization: 'Bearer token'
  - sleep: 10
  - request:
      method: POST
      endpoint: http://example.com/
      body: '{"greet":"Hello World!"}'
      headers:
          content-type: 'application/json'
          Authorization: 'Bearer token'
`)

    input := input{}
    err := yaml.Unmarshal(yamlInput, &input)
    if err != nil {
        log.Fatal(err)
    }

    actions := buildActions(&input)
    return actions
}

func buildActions(input *input) (actions []Action) {
    for _, action := range input.Actions {
        for k, v := range action {
            switch k {
            case "request":
                actionMap := v.(map[interface{}]interface{})
                actions = append(actions,
                    NewRequest(actionMap["method"].(string),
                        actionMap["endpoint"].(string),
                        actionMap["body"].(string),
                        buildHeaders(actionMap["headers"]),
                        getClient(input.Timeout)))
            case "sleep":
                duration := time.Duration(v.(int))
                actions = append(actions, NewSleep(duration*time.Second))
            default:
                fmt.Printf("not much pattern..")
            }
        }
    }
    return
}

func buildHeaders(headers interface{}) map[string]string {
    headersMap := headers.(map[interface{}]interface{})
    headersStrMap := map[string]string{}
    for k, v := range headersMap {
        headersStrMap[k.(string)] = v.(string)
    }
    return headersStrMap
}

func getClient(timeout time.Duration) *http.Client {
    return &http.Client{Timeout: timeout * time.Second}
}
