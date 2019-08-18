package internal

import "log"

type Executor struct {
    dataInput  DataInput
    dataOutput DataOutput
}

func NewExecutor(dataInput DataInput, dataOutput DataOutput) *Executor {
    return &Executor{dataInput: dataInput, dataOutput: dataOutput}
}

func (e *Executor) Execute() {
    actions, err := e.dataInput.LoadActions()
    if err != nil {
        log.Fatal(err)
    }

    for _, action := range actions {
        action.Do()
    }

    e.dataOutput.OutputReport(actions)
}
