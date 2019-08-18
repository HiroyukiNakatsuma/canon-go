package data_output

import (
    "fmt"
    "log"
    "encoding/json"
    "os"

    "github.com/HiroyukiNakatsuma/canon-go/internal/action"
    "github.com/HiroyukiNakatsuma/canon-go/internal/result"
)

type report struct {
    Name      string     `json:"name"`
    Summaries []*summary `json:"summaries"`
}

type summary struct {
    Label               string    `json:"label"`
    ResponseTimeAverage string    `json:"responseTimeAverage"`
    ErrorRate           string    `json:"errorRate"`
    Details             []*detail `json:"details"`
}

type detail struct {
    Status       int    `json:"status"`
    ResponseTime string `json:"responseTime"`
}

type jsonOutput struct {
    Filepath string
}

func NewJsonOutput(filepath string) *jsonOutput {
    if filepath == "" {
        filepath = "./sample/output.json"
    }
    return &jsonOutput{Filepath: filepath}
}

func (output *jsonOutput) OutputReport(actions []action.Action) {
    summaries := output.SummarizeByAction(actions)
    report := report{Name: "Tile", Summaries: summaries}
    bytes, _ := json.Marshal(report)

    file, err := os.Create(output.Filepath)
    if err != nil {
        log.Printf("output error!")
    }
    defer file.Close()

    file.Write(bytes)
}

func (output *jsonOutput) SummarizeByAction(actions []action.Action) (summaries []*summary) {
    for _, act := range actions {
        results := act.GetResults()
        if len(results) == 0 {
            log.Printf("no result.")
            continue
        }

        summaries = append(
            summaries,
            &summary{
                Label:               fmt.Sprintf("%s %s", act.(*action.Request).Method, act.(*action.Request).Url),
                ResponseTimeAverage: fmt.Sprintf("%.2fs", calculateResponseTimeAverage(act.GetResults())),
                ErrorRate:           fmt.Sprintf("%d%%", int(calculateErrorRate(act.GetResults())*100)),
                Details:             mapResult2Detail(act.GetResults()),
            })
    }
    return
}

func mapResult2Detail(results []*result.Result) (details []*detail) {
    for _, r := range results {
        details = append(details, &detail{Status: r.StatusCode, ResponseTime: fmt.Sprintf("%fs", r.ResponseTime.Seconds())})
    }
    return
}

func calculateResponseTimeAverage(results []*result.Result) float32 {
    var num = float64(len(results))
    var sum float64
    for _, r := range results {
        sum = sum + r.ResponseTime.Seconds()
    }
    return float32(sum / num)
}

func calculateErrorRate(results []*result.Result) float64 {
    var num = len(results)
    var errorCount uint
    for _, r := range results {
        if isErrorStatus(r.StatusCode) {
            errorCount += 1
        }
    }
    return float64(errorCount) / float64(num)
}

func isErrorStatus(statusCode int) bool {
    return !(statusCode >= 200 && statusCode <= 399)
}
