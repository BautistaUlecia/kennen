package group

import (
	"errors"
	"fmt"
	"kennen/internal/domain"
)

type CreateGroup struct {
	repository CreateRepository
}

var ErrNameTaken = errors.New("group name already taken")

type CreateRepository interface {
	ExistsByName(name string) (bool, error)
	Save(group *domain.Group) error
}

func NewCreateGroup(r CreateRepository) *CreateGroup {
	return &CreateGroup{repository: r}
}

func (c *CreateGroup) Run(name string) (*domain.Group, error) {
	g, err := domain.NewGroup(name)
	if err != nil {
		return nil, fmt.Errorf("error with group creation: %w", err)
	}

	exists, err := c.repository.ExistsByName(name)
	if err != nil {
		return nil, fmt.Errorf("error finding group by name %w", err)
	}
	if exists {
		return nil, ErrNameTaken
	}

	err = c.repository.Save(g)
	if err != nil {
		return nil, fmt.Errorf("error saving group: %w", err)
	}
	return g, nil
}
