package mock

import "github.com/HiroyukiNakatsuma/canon-go/internal"

type JsonOutputMock struct{}

func NewJsonOutputMock() *JsonOutputMock {
    return &JsonOutputMock{}
}

func (mock *JsonOutputMock) OutputReport(actions []internal.Action) {}
