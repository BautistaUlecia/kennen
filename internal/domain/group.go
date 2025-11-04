package domain

import "fmt"

type Group struct {
	ID        string
	Name      string
	Summoners []Summoner
}

var groupCounter int

func NewGroup(name string) (*Group, error) {
	return &Group{ID: nextGroupID(), Name: name, Summoners: make([]Summoner, 0)}, nil
}

func (g *Group) AddSummoner(s Summoner) {
	g.Summoners = append(g.Summoners, s)
}

func nextGroupID() string {
	groupCounter++
	return fmt.Sprintf("%d", groupCounter)
}
