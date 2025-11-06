package httpgroup

import (
	"fmt"
	"kennen/internal/domain"
)

const profileIconURLTemplate = "https://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%d.png"

type VersionGetter interface {
	GetLatestVersion() string
}

type GroupResponseMapper interface {
	ToGroupResponse(g *domain.Group) GroupResponse
}

type groupResponseMapper struct {
	versionGetter VersionGetter
}

func NewGroupResponseMapper(vg VersionGetter) GroupResponseMapper {
	return &groupResponseMapper{
		versionGetter: vg,
	}
}

func (m *groupResponseMapper) ToGroupResponse(g *domain.Group) GroupResponse {
	out := GroupResponse{ID: g.ID, Name: g.Name}
	out.Summoners = make([]SummonerResponse, 0, len(g.Summoners))
	apiVersion := m.versionGetter.GetLatestVersion()

	for _, s := range g.Summoners {
		iconURL := fmt.Sprintf(profileIconURLTemplate, apiVersion, s.IconID)
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
