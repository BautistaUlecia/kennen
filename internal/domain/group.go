package domain

import "fmt"

type Group struct {
	ID        string
	Name      string
	Summoners []Summoner
	// Tags?
}

var groupCounter int

func NewGroup(name string) (*Group, error) {
	return &Group{ID: nextGroupID(), Name: name, Summoners: make([]Summoner, 0)}, nil
}

func nextGroupID() string {
	groupCounter++
	return fmt.Sprintf("%d", groupCounter)
}
