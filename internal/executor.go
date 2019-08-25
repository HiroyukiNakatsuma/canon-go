package internal

import (
    "log"

    "github.com/HiroyukiNakatsuma/canon-go/internal/data_input"
    "github.com/HiroyukiNakatsuma/canon-go/internal/data_output"
)

type Executor struct {
    dataInput  data_input.DataInput
    dataOutput data_output.DataOutput
}

func NewExecutor(dataInput data_input.DataInput, dataOutput data_output.DataOutput) *Executor {
    return &Executor{dataInput: dataInput, dataOutput: dataOutput}
}

func (e *Executor) Execute() {
    actions := e.dataInput.LoadActions()
    config := e.dataInput.LoadConfig()

    for i := 1; i <= config.Loop; i++ {
        log.Printf("start loop %dth.", i)
        for _, action := range actions {
            action.Do()
        }
    }
    e.dataOutput.OutputReport(actions)
}
