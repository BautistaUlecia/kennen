package infragroup

import "kennen/internal/domain"

type InMemoryRepository struct {
	byName map[string]domain.Group
	byId   map[string]domain.Group
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{byId: make(map[string]domain.Group), byName: make(map[string]domain.Group)}
}

func (r *InMemoryRepository) ExistsByName(name string) (bool, error) {
	_, exists := r.byName[name]
	return exists, nil
}
func (r *InMemoryRepository) Save(g *domain.Group) error {
	r.byName[g.Name] = *g
	r.byId[g.ID] = *g
	return nil
}
