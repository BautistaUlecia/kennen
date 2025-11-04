package group

import (
	"kennen/internal/domain"
)

type ListGroup struct {
	repository ListRepository
}

type ListRepository interface {
	List() ([]*domain.Group, error)
}

func NewListGroup(r ListRepository) *ListGroup {
	return &ListGroup{repository: r}
}

func (g *ListGroup) Run() ([]*domain.Group, error) {
	return g.repository.List()
}
