package internal

type Summary struct {
    Id                  int
    Label               string
    ResponseTimeAverage float32
    ErrorRate           float32
    Results             *[]Result
}
