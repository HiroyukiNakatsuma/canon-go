package internal

import (
    "net/http"
    "strings"
    "time"
)

type YamlLoader struct{}

func NewYamlLoader() *YamlLoader {
    return &YamlLoader{}
}

func (yamlLoader *YamlLoader) LoadConfig() *ActionConfig {
    return &ActionConfig{Threads: 1, Loop: 1}
}

func (yamlLoader *YamlLoader) LoadActions() []Action {
    var actions []Action
    req := NewRequest(
        "POST",
        "http://example.com/",
        `{"greet":"Hello World!"}`,
        BuildHeader("content-type: application/json", "Authorization: Bearer tokenExample"),
        getClient(30))
    actions = append(actions, req)
    return actions
}

/**
Multiple headers define like this:
Accept: application/json,text/csv,application/xml

not like this:
Accept: application/json
Accept: text/csv
Accept: application/xml
*/
func BuildHeader(headers ...string) map[string][]string {
    headersMap := make(http.Header)
    for _, h := range headers {
        h := strings.Split(h, ":")
        k, v := strings.TrimSpace(h[0]), strings.TrimSpace(h[1])
        headersMap[k] = strings.Split(strings.TrimSpace(v), ",")
    }
    return headersMap
}

func getClient(timeout time.Duration) *http.Client {
    return &http.Client{Timeout: timeout * time.Second}
}
