package main

import (
    "flag"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_input"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_output"
)

func main() {
    var inputFilepath = flag.String("input", "./sample/input.yml", "input file path.")
    var outputFilepath = flag.String("output", "./sample/output.json", "output file path.")
    flag.Parse()
    internal.NewExecutor(data_input.NewYamlLoader(*inputFilepath), data_output.NewJsonOutput(*outputFilepath)).Execute()
}
