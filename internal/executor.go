package internal

type Executor struct {
    dataInput  DataInput
    dataOutput DataOutput
}

func NewExecutor(dataInput DataInput, dataOutput DataOutput) *Executor {
    return &Executor{dataInput: dataInput, dataOutput: dataOutput}
}

func (e *Executor) Execute() {
    actions := e.dataInput.LoadActions()

    for _, action := range actions {
        action.Do()
    }

    e.dataOutput.OutputReport(actions)
}
