package main

import "github.com/HiroyukiNakatsuma/canon-go/internal"

func main() {
    internal.NewExecutor(internal.NewYamlLoader(), internal.NewJsonOutput()).Execute()
}
