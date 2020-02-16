package port

import "kubetui/domain"

type ContextPort interface {
	FindContexts() domain.Contexts
}
