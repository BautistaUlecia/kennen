package group

import (
	"errors"
	"kennen/internal/domain"
)

var ErrNotFound = errors.New("group not found")

type GetGroup struct {
	repository GetRepository
}

type GetRepository interface {
	GetByID(ID string) (*domain.Group, error)
}

func NewGetGroup(r GetRepository) *GetGroup {
	return &GetGroup{repository: r}
}

func (g *GetGroup) Run(ID string) (*domain.Group, error) {
	return g.repository.GetByID(ID)
}
