package internal

import (
    "log"
    "io/ioutil"
)

func Execute(req *Request) {
    log.Printf("req: %v", req)

    api := NewApi(req, nil)
    res, err := api.DoRequest()
    if err != nil {
        log.Fatal(err)
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Response Status: %d", res.StatusCode)
    log.Printf("Response Body: %s", b)
}
