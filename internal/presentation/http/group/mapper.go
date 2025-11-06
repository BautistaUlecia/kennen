package httpgroup

import (
	"fmt"
	"kennen/internal/domain"
)

type VersionGetter interface {
	GetLatestVersion() string
}

func toGroupResponse(g *domain.Group, vg VersionGetter) GroupResponse {
	out := GroupResponse{ID: g.ID, Name: g.Name}
	out.Summoners = make([]SummonerResponse, 0, len(g.Summoners))
	apiVersion := vg.GetLatestVersion()

	for _, s := range g.Summoners {
		iconURL := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%d.png", apiVersion, s.IconID)
		out.Summoners = append(out.Summoners, SummonerResponse{
			Name:    s.Name,
			Tier:    s.Tier,
			Rank:    s.Rank,
			LP:      s.LeaguePoints,
			Wins:    s.Wins,
			Losses:  s.Losses,
			IconURL: iconURL,
			Level:   s.Level,
		})
	}
	return out
}
