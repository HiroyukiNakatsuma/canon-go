package main

import "github.com/HiroyukiNakatsuma/canon-go/internal"

func main() {
    internal.NewExecutor(internal.NewYamlLoader(), internal.NewSummarizer(), internal.NewJsonOutput()).Execute()
}
