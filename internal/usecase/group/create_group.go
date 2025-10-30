package group

import (
	"fmt"
	"kennen/internal/domain"
)

type CreateGroup struct {
	repository Repository
}
type Repository interface {
	ExistsByName(name string) (bool, error)
	Save(group *domain.Group) error
}

func NewCreateGroup(repository Repository) *CreateGroup {
	return &CreateGroup{repository: repository}
}

func (c *CreateGroup) Run(name string) error {
	g, err := domain.NewGroup(name)
	if err != nil {
		return fmt.Errorf("error with group creation: %w", err)
	}

	exists, err := c.repository.ExistsByName(name)
	if err != nil {
		return fmt.Errorf("error finding group by name %w", err)
	}
	if exists {
		return fmt.Errorf("group name is already taken")
	}

	err = c.repository.Save(g)
	if err != nil {
		return fmt.Errorf("error saving group: %w", err)
	}
	return nil
}
