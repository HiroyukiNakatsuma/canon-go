package internal

import "log"

func Execute(req *Request) {
    log.Printf("req: %v", req)

    api := API{Req: req}
    res := api.DoRequest()

    log.Print(res)
}
