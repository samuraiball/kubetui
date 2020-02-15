package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"kubetui/domain"
	"kubetui/usecase"
	"testing"
)

func Test指定したネームスペースのPodの一覧を取得することができる(t *testing.T) {

	podGatewayMock := &PodUsecaseMock{}
	podGatewayMock.On("FindPods", "").Return(
		domain.Pods{
			Pods: []domain.Pod{
				{
					Name:   "sample-deployment-7bb5fc6bc6-9btmd",
					Ip:     "172.17.0.17",
					Status: "Running",
					Node:   "minikube",
				},
				{
					Name:   "sample-deployment-7bb5fc6bc6-hhd2l",
					Ip:     "172.17.0.19",
					Status: "Running",
					Node:   "minikube",
				},
			},
		},
	)
	target := usecase.PodUsecase{PodPort: podGatewayMock}

	expected := []string{"sample-deployment-7bb5fc6bc6-9btmd", "sample-deployment-7bb5fc6bc6-hhd2l"}

	actual := target.GetPodList("")

	assert.Equal(t, expected, actual)
}

func (p PodUsecaseMock) FindPods(namespace string) domain.Pods {
	result := p.Called(namespace).Get(0)
	return result.(domain.Pods)
}
