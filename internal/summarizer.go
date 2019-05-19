package internal

type Summarizer interface {
    Summarize()
}

type summarizer struct{}

func NewSummarizer() *summarizer {
    return &summarizer{}
}

func (summarizer *summarizer) Summarize() {}
