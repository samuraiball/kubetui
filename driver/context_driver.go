package driver

import "os/exec"

type ContextDriver interface {
	FindContexts() []byte
}

type ContextK8sDriver struct {
}

func (c ContextK8sDriver) FindContexts() []byte {
	contextJsonByte, err := exec.Command("kubectl", "config", "get-contexts", "-o", "name").Output()
	if err != nil {
		panic(err)
	}
	return contextJsonByte
}
