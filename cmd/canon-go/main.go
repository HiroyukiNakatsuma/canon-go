package main

import (
    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "log"
)

func main() {
    body := `{"greet":"Hello World!"}`
    httpMethod := "POST"
    endpoint := "http://example.com/"
    header1 := "content-type: application/json"
    header2 := "Authorization: Bearer tokenExample"
    var req = internal.Request{Method: httpMethod, Endpoint: endpoint, Body: body, Headers: internal.BuildHeader(header1, header2)}
    log.Printf("req: %v", req)

    api := internal.API{Req: &req}
    res := api.Execute()
    log.Print(res)
}
