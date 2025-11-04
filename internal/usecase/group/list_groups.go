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

func NewListGroup(repository ListRepository) *ListGroup {
	return &ListGroup{repository: repository}
}

func (g *ListGroup) Run() ([]*domain.Group, error) {
	return g.repository.List()
}
