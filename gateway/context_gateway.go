package gateway

import (
	"kubetui/domain"
	"kubetui/driver"
	"strings"
)

type ContextGateway struct {
	ContextDriver driver.ContextDriver
}

func (c ContextGateway) FindContexts() domain.Contexts {
	rowData := c.ContextDriver.FindContexts()
	stringData := string(rowData)
	contextList := strings.Split(stringData, " ")

	contexts := make([]domain.Context, len(contextList))
	for i, contextString := range contextList {
		contexts[i] = domain.Context{Name: contextString}
	}

	return domain.Contexts{Contexts: contexts}
}
