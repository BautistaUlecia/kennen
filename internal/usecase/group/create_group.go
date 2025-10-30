package usecase

type CreateGroup struct {
}

func (c *CreateGroup) Run() {

}

func NewCreateGroup() *CreateGroup {
	return &CreateGroup{}
}
