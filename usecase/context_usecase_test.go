package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"kubetui/domain"
	"kubetui/usecase"
	"testing"
)

func TestContextの一覧を取得する(t *testing.T) {

	contextGatewayMock := &ContextGatewayMock{}
	contextGatewayMock.On("FindContexts").Return(
		domain.Contexts{
			Contexts: []domain.Context{
				{
					Name: "minikube",
				},
				{
					Name: "testContext",
				},
			}})
	target := usecase.ContextUsecase{ContextPort: contextGatewayMock}

	actual := target.GetContextList()
	expected := []string{"minikube", "testContext"}

	assert.Equal(t, expected, actual)
}

func (c ContextGatewayMock) FindContexts() domain.Contexts {
	result := c.Called().Get(0)
	return result.(domain.Contexts)
}
