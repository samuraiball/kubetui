package port

import "kubetui/domain"

type PodPort interface {
	FindPods(namespace string) domain.Pods
}
