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

func Execute() {
    var client = getClient()
    body := `{"greet":"Hello World!"}`
    httpMethod := "POST"
    endpoint := "http://example.com/"
    log.Printf("Start %s %s", httpMethod, endpoint)
    log.Printf("Params: %s", body)
    req, err := http.NewRequest(httpMethod, endpoint, bytes.NewBuffer([]byte(body)))
    req.Header.Add("content-type", "application/json")
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
