package internal

import (
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
    for _, action := range actions {
        action.Do()
    }
    e.dataOutput.OutputReport(actions)
}
