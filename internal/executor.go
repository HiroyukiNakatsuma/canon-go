package internal

import "log"

type Executor struct {
    Actions []Action
}

func NewExecutor(actions []Action) *Executor {
    return &Executor{Actions: actions}
}

func (e *Executor) Do() {
    for _, action := range e.Actions {
        log.Printf("req: %v", action)

        result := action.Do()

        log.Printf("Response Status: %d", result.StatusCode)
        log.Printf("Response Body: %s", result.Body)
        log.Printf("Response Time: %fs", result.Time.Seconds())
    }
}
