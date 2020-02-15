package gateway

import (
	"encoding/json"
	"kubetui/domain"
	"kubetui/driver"
)

type PodGateway struct {
	PodDriver driver.PodDriver
}

func (g PodGateway) FindPods(namespace string) domain.Pods {

	rowData := g.PodDriver.FindPods(namespace)

	var podJsonStruct PodJsonStruct
	if err := json.Unmarshal(rowData, &podJsonStruct); err != nil {
		panic(err)
	}

	podList := make([]domain.Pod, len(podJsonStruct.Items))

	for i, data := range podJsonStruct.Items {
		podList[i] = domain.Pod{
			Name:   data.Metadata.Name,
			Ip:     data.Status.PodIP,
			Status: data.Status.Phase,
			Node:   data.Spec.NodeName,
		}
	}
	pods := domain.Pods{Pods: podList}

	return pods
}

type PodJsonStruct struct {
	Items []Item `json:"items"`
}
type Item struct {
	Metadata Metadata `json:"metadata"`
	Status   Status   `json:"status"`
	Spec     Spec     `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	NodeName string `json:"nodeName"`
}

type Status struct {
	PodIP string `json:"podIP"`
	Phase string `json:"phase"`
}
