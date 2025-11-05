package httpgroup

type AddToGroupRequest struct {
	Region   string `json:"region"`
	GameName string `json:"game_name"`
	Tag      string `json:"tag"`
}
type CreateGroupRequest struct {
	Name string `json:"name"`
}
