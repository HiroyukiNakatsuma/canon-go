package main

import (
    "flag"

    "github.com/HiroyukiNakatsuma/canon-go/internal"
)

func main() {
    var inputFilepath = flag.String("input", "./sample/input.yml", "input file path.")
    flag.Parse()
    internal.NewExecutor(internal.NewYamlLoader(*inputFilepath), internal.NewJsonOutput()).Execute()
}
