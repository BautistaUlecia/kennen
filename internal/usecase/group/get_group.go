package usecase

type GetGroup struct {
}

func (g *GetGroup) Run() {

}

func NewGetGroup() *GetGroup {
	return &GetGroup{}
}
