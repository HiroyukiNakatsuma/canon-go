package data_input

import (
    "github.com/HiroyukiNakatsuma/canon-go/internal/config"
    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
)

type DataInput interface {
    LoadConfig() *config.ActionConfig
    LoadActions() ([]action.Action, error)
}
