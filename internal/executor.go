package internal

import "log"

type Executor struct {
    dataInput  DataInput
    summarizer Summarizer
}

func NewExecutor(dataInput DataInput, summarizer Summarizer) *Executor {
    return &Executor{dataInput: dataInput, summarizer: summarizer}
}

func (e *Executor) Execute() {
    actions := e.dataInput.LoadActions()
    for _, action := range actions {
        log.Printf("req: %v", action)

        result := action.Do()

        if result == nil {
            log.Printf("finish sleep...")
            continue
        }
        log.Printf("Response Status: %d", result.StatusCode)
        log.Printf("Response Body: %s", result.Body)
        log.Printf("Response Time: %fs", result.Time.Seconds())
    }
}
