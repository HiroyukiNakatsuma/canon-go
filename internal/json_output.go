package internal

import (
    "fmt"
    "log"
    "encoding/json"
    "os"
)

type report struct {
    name      string
    summaries []*summary
}

type summary struct {
    id                  int
    label               string
    responseTimeAverage string
    errorRate           string
    details             []*detail
}

type detail struct {
    status       int
    responseTime string
}

type jsonOutput struct{}

func NewJsonOutput() *jsonOutput {
    return &jsonOutput{}
}

func (output *jsonOutput) OutputReport(actions []Action) {
    summaries := Summarize(actions)
    report := report{name: "Tile", summaries: summaries}
    bytes, _ := json.Marshal(report)

    file, err := os.Create(`./report.json`)
    if err != nil {
        log.Printf("output error!")
    }
    defer file.Close()

    file.Write(bytes)
}

func Summarize(actions []Action) (summaries []*summary) {
    for i, action := range actions {
        results := action.GetResults()
        if len(results) == 0 {
            log.Printf("no result.")
            continue
        }

        summaries = append(
            summaries,
            &summary{
                id:                  i,
                label:               fmt.Sprintf("%s %s", results[0].Request.Method, results[0].Request.Endpoint),
                responseTimeAverage: fmt.Sprintf("%f", calculateResponseTimeAverage(action.GetResults())),
                errorRate:           fmt.Sprintf("%f", calculateErrorRate(action.GetResults())),
                details:             mapResult2Detail(action.GetResults()),
            })
    }
    return
}

func mapResult2Detail(results []*Result) (details []*detail) {
    for _, result := range results {
        details = append(details, &detail{status: result.StatusCode, responseTime: fmt.Sprintf("%fs", result.ResponseTime.Seconds())})
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
