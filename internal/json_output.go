package internal

import (
    "fmt"
    "log"
    "encoding/json"
    "os"
)

type report struct {
    Name      string     `json:"name"`
    Summaries []*summary `json:"summaries"`
}

type summary struct {
    Id                  int       `json:"id"`
    Label               string    `json:"label"`
    ResponseTimeAverage string    `json:"responseTimeAverage"`
    ErrorRate           string    `json:"errorRate"`
    Details             []*detail `json:"details"`
}

type detail struct {
    Status       int    `json:"status"`
    ResponseTime string `json:"responseTime"`
}

type jsonOutput struct{}

func NewJsonOutput() *jsonOutput {
    return &jsonOutput{}
}

func (output *jsonOutput) OutputReport(actions []Action) {
    summaries := summarizeByAction(actions)
    report := report{Name: "Tile", Summaries: summaries}
    bytes, _ := json.Marshal(report)

    file, err := os.Create(`./report.json`)
    if err != nil {
        log.Printf("output error!")
    }
    defer file.Close()

    file.Write(bytes)
}

func summarizeByAction(actions []Action) (summaries []*summary) {
    for i, action := range actions {
        results := action.GetResults()
        if len(results) == 0 {
            log.Printf("no result.")
            continue
        }

        summaries = append(
            summaries,
            &summary{
                Id:                  i,
                Label:               fmt.Sprintf("%s %s", results[0].Request.Method, results[0].Request.Endpoint),
                ResponseTimeAverage: fmt.Sprintf("%f", calculateResponseTimeAverage(action.GetResults())),
                ErrorRate:           fmt.Sprintf("%f", calculateErrorRate(action.GetResults())),
                Details:             mapResult2Detail(action.GetResults()),
            })
    }
    return
}

func mapResult2Detail(results []*Result) (details []*detail) {
    for _, result := range results {
        details = append(details, &detail{Status: result.StatusCode, ResponseTime: fmt.Sprintf("%fs", result.ResponseTime.Seconds())})
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
