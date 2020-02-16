package usecase

import "kubetui/port"

type ContextUsecase struct {
	ContextPort port.ContextPort
}

func (c ContextUsecase) GetContextList() []string {
	contexts := c.ContextPort.FindContexts()
	contextList := make([]string, len(contexts.Contexts))
	for i, context := range contexts.Contexts {
		contextList[i] = context.Name
	}
	return contextList
}
