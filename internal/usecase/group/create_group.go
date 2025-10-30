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

func NewCreateGroup(repository CreateRepository) *CreateGroup {
	return &CreateGroup{repository: repository}
}

func (c *CreateGroup) Run(name string) (string, error) {
	g, err := domain.NewGroup(name)
	if err != nil {
		return "", fmt.Errorf("error with group creation: %w", err)
	}

	exists, err := c.repository.ExistsByName(name)
	if err != nil {
		return "", fmt.Errorf("error finding group by name %w", err)
	}
	if exists {
		return "", ErrNameTaken
	}

	err = c.repository.Save(g)
	if err != nil {
		return "", fmt.Errorf("error saving group: %w", err)
	}
	return g.ID, nil
}
