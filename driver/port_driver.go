package driver

import "os/exec"

type PodDriver interface {
	FindPods(namespace string) []byte
}

type PodK8sDriver struct {
}

func (ps PodK8sDriver) FindPods(namespase string) []byte {
	podJsonByte, err := exec.Command("kubectl", "-n", namespase, "get", "po", "-o", "json").Output()
	if err != nil {
		panic(err)
	}
	return podJsonByte
}
