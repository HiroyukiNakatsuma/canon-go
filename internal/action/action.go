package action

import "github.com/HiroyukiNakatsuma/canon-go/internal/result"

type Action interface {
    Do()
    GetResults() []*result.Result
}
