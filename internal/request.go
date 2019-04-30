package internal

import (
    "strings"
    "net/http"
)

/**
Multiple headers define like this:
Accept: application/json,text/csv,application/xml

not like this:
Accept: application/json
Accept: text/csv
Accept: application/xml
*/
type Request struct {
    Method   string
    Endpoint string
    Body     string
    Headers  map[string][]string
}

func BuildHeader(headers ...string) map[string][]string {
    headersMap := make(http.Header)
    for _, h := range headers {
        h := strings.Split(h, ":")
        k, v := strings.TrimSpace(h[0]), strings.TrimSpace(h[1])
        headersMap[k] = strings.Split(strings.TrimSpace(v), ",")
    }
    return headersMap
}
