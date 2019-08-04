package internal

import (
    "net/http"
    "time"
    "log"
    "fmt"
    "io/ioutil"
    "errors"

    "gopkg.in/yaml.v2"
)

type yamlLoader struct {
    Filepath string
}

type input struct {
    Timeout time.Duration
    Actions []map[string]interface{}
}

func NewYamlLoader(filepath string) *yamlLoader {
    if filepath == "" {
        filepath = "./sample/sample.yml"
    }
    return &yamlLoader{Filepath: filepath}
}

func (yamlLoader *yamlLoader) LoadConfig() *ActionConfig {
    return &ActionConfig{Threads: 1, Loop: 1}
}

func (yamlLoader *yamlLoader) LoadActions() []Action {
    yamlInput, err := ioutil.ReadFile(yamlLoader.Filepath)
    if err != nil {
        panic(err)
    }

    input := input{}
    err = yaml.Unmarshal(yamlInput, &input)
    if err != nil {
        log.Fatal(err)
    }

    actions, err := buildActions(&input)
    if err != nil {
        log.Fatal(err)
    }

    return actions
}

func buildActions(input *input) (actions []Action, err error) {
    for _, action := range input.Actions {
        for k, v := range action {
            switch k {
            case "request":
                actionMap := v.(map[interface{}]interface{})
                req, err := buildRequest(actionMap, input.Timeout)
                if err != nil {
                    return nil, err
                }
                actions = append(actions, req)
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

func buildRequest(requestMap map[interface{}]interface{}, timeout time.Duration) (req *Request, error error) {
    method := requestMap["method"]
    url := requestMap["url"]
    body := requestMap["body"]
    if method == nil {
        return nil, errors.New("method is required.")
    }
    if url == nil {
        return nil, errors.New("url is required.")
    }
    if body == nil {
        body = ""
    }

    var headers map[string]string
    if requestMap["headers"] != nil {
        headers = buildHeaders(requestMap["headers"])
    }

    return NewRequest(method.(string), url.(string), body.(string), headers, getClient(timeout)), nil
}

func buildHeaders(headers interface{}) map[string]string {
    headersMap := headers.(map[interface{}]interface{})
    headersStrMap := make(map[string]string)
    for k, v := range headersMap {
        headersStrMap[k.(string)] = v.(string)
    }
    return headersStrMap
}

func getClient(timeout time.Duration) *http.Client {
    return &http.Client{Timeout: timeout * time.Second}
}
