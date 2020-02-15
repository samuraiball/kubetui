package driver

import "os/exec"

type PodDriver interface {
	FindPods(namespace string) []byte
}

type K8sPodDriver struct {
}

func (ps K8sPodDriver) FindPods(namespase string) []byte {
	podJsonByte, err := exec.Command("kubectl", "-n", namespase, "get", "po", "-o", "json").Output()
	if err != nil {
		panic(err)
	}
	return podJsonByte
}
