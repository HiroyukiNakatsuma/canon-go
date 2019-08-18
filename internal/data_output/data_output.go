package data_output

import "github.com/HiroyukiNakatsuma/canon-go/internal/action"

type DataOutput interface {
    OutputReport(actions []action.Action)
}
