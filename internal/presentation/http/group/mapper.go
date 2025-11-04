package httpgroup

import "kennen/internal/domain"

func toGroupResponse(g *domain.Group) GroupResponse {
	out := GroupResponse{ID: g.ID, Name: g.Name}
	out.Summoners = make([]SummonerResponse, 0, len(g.Summoners))
	for _, s := range g.Summoners {
		out.Summoners = append(out.Summoners, SummonerResponse{
			Name: s.Name, Tier: s.Tier, Rank: s.Rank, LP: s.LeaguePoints, Wins: s.Wins, Losses: s.Losses,
		})
	}
	return out
}
