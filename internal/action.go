package internal

type Action interface {
    Do()
    GetResults() []*Result
}
