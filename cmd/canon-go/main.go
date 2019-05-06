package main

import "github.com/HiroyukiNakatsuma/canon-go/internal"

func main() {
    var actions []internal.Action

    body := `{"greet":"Hello World!"}`
    httpMethod := "POST"
    endpoint := "http://example.com/"
    header1 := "content-type: application/json"
    header2 := "Authorization: Bearer tokenExample"
    var req = &internal.Request{
        Method:   httpMethod,
        Endpoint: endpoint,
        Body:     body,
        Headers:  internal.BuildHeader(header1, header2),
        Timeout:  30,
    }
    actions = append(actions, req)
    internal.NewExecutor(actions).Do()
}
