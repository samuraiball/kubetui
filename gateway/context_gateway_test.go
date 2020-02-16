package gateway_test

import (
	"github.com/stretchr/testify/assert"
	"kubetui/domain"
	"kubetui/gateway"
	"testing"
)

func TestContextの一覧を取得することができる(t *testing.T) {

	contextDriverMock := &ContextDriverMock{}
	contextDriverMock.On("FindContexts").Return([]byte("minikube textContext"))

	target := gateway.ContextGateway{
		ContextDriver: contextDriverMock,
	}

	expected := domain.Contexts{Contexts: []domain.Context{
		{
			Name: "minikube",
		},
		{
			Name: "textContext",
		},
	}}

	actual := target.FindContexts()
	assert.Equal(t, expected, actual)
}

func (c *ContextDriverMock) FindContexts() []byte {
	result := c.Called().Get(0)
	return result.([]byte)
}
