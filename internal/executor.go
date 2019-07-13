package internal

import "log"

type Executor struct {
    dataInput  DataInput
    summarizer Summarizer
    dataOutput DataOutput
}

func NewExecutor(dataInput DataInput, summarizer Summarizer, dataOutput DataOutput) *Executor {
    return &Executor{dataInput: dataInput, summarizer: summarizer, dataOutput: dataOutput}
}

func (e *Executor) Execute() {
    actions := e.dataInput.LoadActions()
    var results []*Result
    for _, action := range actions {
        log.Printf("req: %v", action)

        result := action.Do()

        if result == nil {
            log.Printf("finish sleep...")
            continue
        }
        log.Printf("Response Status: %d", result.StatusCode)
        log.Printf("Response Body: %s", result.ResponseBody)
        log.Printf("Response Time: %fs", result.ResponseTime.Seconds())

        results = append(results, result)
    }

    summary := e.summarizer.Summarize(results)

    e.dataOutput.OutputSummary(summary)
}
