package internal

import (
    "io/ioutil"
    "net/http"
    "bytes"
    "time"
    "fmt"
    "log"
)

func getClient() *http.Client {
    return &http.Client{Timeout: 30 * time.Second}
}

func Execute(request *Request) {
    log.Printf("Request: %v", request)

    var client = getClient()
    req, err := http.NewRequest(request.Method, request.Endpoint, bytes.NewBuffer([]byte(request.Body)))
    for k, v := range request.Headers {
        req.Header.Add(k, v)
    }
    res, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    b, err := ioutil.ReadAll(res.Body)
    if err == nil {
        fmt.Println(string(b))
    }

    log.Printf("Response Status: %d", res.StatusCode)
    log.Printf("Response Body: %s", b)
}
