package internal

type jsonOutput struct{}

func NewJsonOutput() *jsonOutput {
    return &jsonOutput{}
}

func (output *jsonOutput) OutputSummary(summary []*Summary) {
}
