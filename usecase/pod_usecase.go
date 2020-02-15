package usecase

import (
	"kubetui/port"
)

type PodUsecase struct {
	PodPort port.PodPort
}

func (p PodUsecase) GetPodList(namespace string) []string {
	pods := p.PodPort.FindPods(namespace)
	return pods.List()
}

//func (p PodUsecase) describe(targetPodName string, namespace string) string {
//	pods := p.findPods(namespace)
//	return pods.DescribeTargetPod(targetPodName)
//}
