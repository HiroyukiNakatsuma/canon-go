package internal

import (
    "log"
    "fmt"
)

type Summarizer interface {
    Summarize(results []Action) []*Summary
}

type summarizer struct{}

func NewSummarizer() *summarizer {
    return &summarizer{}
}

func (summarizer *summarizer) Summarize(actions []Action) (summaries []*Summary) {
    for i, action := range actions {
        results := action.GetResults()
        if len(results) == 0 {
            log.Printf("no result.")
            continue
        }

        summaries = append(
            summaries,
            &Summary{
                Id:                  i,
                Label:               fmt.Sprintf("%s %s", results[0].Request.Method, results[0].Request.Endpoint),
                ResponseTimeAverage: calculateResponseTimeAverage(action.GetResults()),
                ErrorRate:           calculateErrorRate(action.GetResults()),
                Results:             action.GetResults(),
            })
    }
    return
}

func calculateResponseTimeAverage(results []*Result) float32 {
    var num = float64(len(results))
    var sum float64
    for _, result := range results {
        sum = sum + result.ResponseTime.Seconds()
    }
    return float32(sum / num)
}

func calculateErrorRate(results []*Result) float32 {
    var num = float32(len(results))
    var errorCount uint
    for _, result := range results {
        if isErrorStatus(result.StatusCode) {
            errorCount += 1
        }
    }
    return float32(errorCount) / num
}

func isErrorStatus(statusCode int) bool {
    return !(statusCode >= 200 && statusCode <= 399)
}
