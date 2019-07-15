package internal

type DataOutput interface {
    OutputReport(actions []Action)
}
