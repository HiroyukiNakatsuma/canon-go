package internal

import "strings"

type Request struct {
    Method   string
    Endpoint string
    Body     string
    Headers  map[string]string
}

func BuildHeader(headers ...string) map[string]string {
    headersMap := make(map[string]string)
    for _, h := range headers {
        h := strings.Split(h, ":")
        k, v := strings.TrimSpace(h[0]), strings.TrimSpace(h[1])
        headersMap[k] = strings.TrimSpace(v)
    }
    return headersMap
}
