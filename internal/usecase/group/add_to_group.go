package group

import (
	"fmt"
	"kennen/internal/domain"
	"kennen/internal/infrastructure/riot"
)

type AddToGroup struct {
	repository AddToRepository
	riotClient *riot.Client
}
type AddToRepository interface {
	GetByID(ID string) (*domain.Group, error)
	Save(group *domain.Group) error
}

func NewAddToGroup(r AddToRepository, c *riot.Client) *AddToGroup {
	return &AddToGroup{repository: r, riotClient: c}
}

func (a *AddToGroup) Run(ID, gameName, tag, region string) error {
	g, err := a.repository.GetByID(ID)
	if err != nil {
		return fmt.Errorf("error getting group by id: %w", err)
	}
	s, err := a.riotClient.FindSummoner(gameName, tag, region)
	if err != nil {
		return fmt.Errorf("error getting summoner by name: %v, %v. %w", gameName, tag, err)
	}

	g.AddSummoner(*s)

	err = a.repository.Save(g)
	if err != nil {
		return fmt.Errorf("error saving group: %w", err)
	}
	return nil
}
