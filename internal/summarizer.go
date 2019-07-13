package internal

type Summarizer interface {
    Summarize(results []*Result) *Summary
}

type summarizer struct{}

func NewSummarizer() *summarizer {
    return &summarizer{}
}

func (summarizer *summarizer) Summarize(results []*Result) *Summary {
    return &Summary{}
}
