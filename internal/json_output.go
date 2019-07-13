package internal

type JsonOutput struct{}

func NewJsonOutput() *JsonOutput {
    return &JsonOutput{}
}

func (output *JsonOutput) OutputSummary(summary *Summary) {
}
