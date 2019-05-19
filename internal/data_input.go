package internal

type DataInput interface {
    LoadConfig() *ActionConfig
    LoadActions() []Action
}
