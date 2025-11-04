package httpgroup

type GroupResponse struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Summoners []SummonerResponse `json:"summoners"`
}
type SummonerResponse struct {
	Name   string `json:"name"`
	Tier   string `json:"tier"`
	Rank   string `json:"rank"`
	LP     int    `json:"lp"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
}
