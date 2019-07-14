package internal

type Summarizer interface {
    Summarize(results []Action) *Summary
}

type summarizer struct{}

func NewSummarizer() *summarizer {
    return &summarizer{}
}

func (summarizer *summarizer) Summarize(results []Action) *Summary {
    return &Summary{}
}
