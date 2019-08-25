package data_input

import (
    "net/http"
    "time"
    "log"
    "io/ioutil"
    "errors"

    "gopkg.in/yaml.v2"

    "github.com/HiroyukiNakatsuma/canon-go/internal/config"
    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
)

type yamlLoader struct {
    actions []action.Action
    config  *config.Config
}

type input struct {
    Loop    int
    Timeout time.Duration
    Actions []map[string]interface{}
}

func NewYamlLoader(filepath string) (yl *yamlLoader, err error) {
    if filepath == "" {
        filepath = "./sample/input.yml"
    }

    yamlInput, err := ioutil.ReadFile(filepath)
    if err != nil {
        return nil, errors.New("invalid filepath")
    }

    input := input{}
    err = yaml.Unmarshal(yamlInput, &input)
    if err != nil {
        log.Fatal(err)
        return nil, errors.New("invalid input")
    }

    actions, err := buildActions(&input)
    if err != nil {
        return nil, err
    }

    if input.Loop < 1 {
        input.Loop = 1
    }
    return &yamlLoader{actions: actions, config: &config.Config{Threads: 1, Loop: input.Loop}}, err
}

func (yamlLoader *yamlLoader) LoadConfig() *config.Config {
    return yamlLoader.config
}

func (yamlLoader *yamlLoader) LoadActions() []action.Action {
    return yamlLoader.actions
}

func buildActions(input *input) (actions []action.Action, err error) {
    for _, act := range input.Actions {
        for k, v := range act {
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
                actions = append(actions, action.NewSleep(duration))
            default:
                return nil, errors.New("invalid input")
            }
        }
    }
    return
}

func buildRequest(requestMap map[interface{}]interface{}, timeout time.Duration) (req *action.Request, err error) {
    method := requestMap["method"]
    url := requestMap["url"]
    body := requestMap["body"]
    if method == nil {
        return nil, errors.New("method is required")
    }
    if url == nil {
        return nil, errors.New("url is required")
    }
    if body == nil {
        body = ""
    }

    var headers map[string]string
    if requestMap["headers"] != nil {
        headers = buildHeaders(requestMap["headers"])
    }

    return action.NewRequest(method.(string), url.(string), body.(string), headers, getClient(timeout)), nil
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
