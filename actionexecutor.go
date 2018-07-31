package arc

import (
	"fmt"
)

type ActionExecutor interface {
	GetName() string
	Execute(request ActionRequest) ActionResponse
}

var registeredExecutors []ActionExecutor

func RegisterExecutor(executors ...ActionExecutor) {
	registeredExecutors = append(executors, executors...)
}

func lookupActionExecutor(actionName string) (ActionExecutor, error) {
	for _, executor := range registeredExecutors {
		if executor.GetName() == actionName {
			return executor, nil
		}
	}
	return nil, fmt.Errorf("could not find action executor for : %s", actionName)
}
