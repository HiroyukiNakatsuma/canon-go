package internal

type Action interface {
    Do() *Result
}
