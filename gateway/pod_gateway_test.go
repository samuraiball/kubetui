package gateway_test

import (
	"github.com/stretchr/testify/assert"
	"kubetui/domain"
	"kubetui/gateway"
	"testing"
)

func Test指定されたネームスペースのPodsを全件取得する(t *testing.T) {

	expected := domain.Pods{
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
	}

	podDriverMock := &PodDriverMock{}
	podDriverMock.On("FindPods", "").Return(
		[]byte(podRowData),
	)
	podGateway := gateway.PodGateway{
		PodDriver: podDriverMock,
	}
	actual := podGateway.FindPods("")
	assert.Equal(t, expected, actual)
}

func (p *PodDriverMock) FindPods(namespace string) []byte {
	result := p.Called(namespace).Get(0)
	return result.([]byte)
}
