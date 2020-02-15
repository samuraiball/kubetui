package domain

type Pod struct {
	Name   string
	Ip     string
	Status string
	Node   string
	Image  string
}

type Pods struct {
	Pods []Pod
}

func (p *Pods) List() []string {
	podNames := make([]string, len(p.Pods))
	for i, pod := range p.Pods {
		podNames[i] = pod.Name
	}
	return podNames
}

//func (p *Pods) DescribeTargetPod(targetPodName string) string {
//
//	var targetPodDescription string
//	for _, pod := range p.Pods {
//		if pod.Name == targetPodName {
//			targetPodDescription = pod.Description
//		}
//	}
//
//	if targetPodDescription == "" {
//		panic("Pod Not Found")
//	}
//	return targetPodDescription
//}
