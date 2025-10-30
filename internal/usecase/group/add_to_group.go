package usecase

type AddToGroup struct {
}

func (a *AddToGroup) Run() {}

func NewAddToGroup() *AddToGroup {
	return &AddToGroup{}
}
