package group

import "kennen/internal/domain"

type CreateGroup struct {
}

var counter int = 0

func (c *CreateGroup) Run(name string) {
	domain.NewGroup(name)
}

func NewCreateGroup() *CreateGroup {
	return &CreateGroup{}
}
