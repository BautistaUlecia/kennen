package routine

import (
	"kennen/internal/domain"
	"kennen/internal/infrastructure/riot"
	"log"
	"strings"
	"time"
)

type EloUpdater struct {
	riotClient *riot.Client
	repository GroupRepository
}

type GroupRepository interface {
	List() ([]*domain.Group, error)
	Save(g *domain.Group) error
}

func NewEloUpdater(riotClient *riot.Client, repository GroupRepository) *EloUpdater {
	return &EloUpdater{
		riotClient: riotClient,
		repository: repository,
	}
}

func (eu *EloUpdater) Start() {
	log.Println("Starting initial elo update...")
	if err := eu.updateAllSummoners(); err != nil {
		log.Printf("Error in initial elo update: %v", err)
	}

	ticker := time.NewTicker(15 * time.Minute)
	go func() {
		for range ticker.C {
			log.Println("Starting scheduled elo update...")
			if err := eu.updateAllSummoners(); err != nil {
				log.Printf("Error in scheduled elo update: %v", err)
			}
		}
	}()
}

func (eu *EloUpdater) updateAllSummoners() error {
	groups, err := eu.repository.List()
	if err != nil {
		return err
	}

	updatedCount := 0
	errorCount := 0

	for _, group := range groups {
		for i := range group.Summoners {
			summoner := &group.Summoners[i]

			parts := strings.Split(summoner.Name, "#")
			if len(parts) != 2 {
				log.Printf("Invalid summoner name format: %s", summoner.Name)
				errorCount++
				continue
			}

			gameName := parts[0]
			tagLine := parts[1]

			updatedSummoner, err := eu.riotClient.FindSummoner(gameName, tagLine, "la2")
			if err != nil {
				log.Printf("Error updating summoner %s: %v", summoner.Name, err)
				errorCount++
				continue
			}

			summoner.Tier = updatedSummoner.Tier
			summoner.Rank = updatedSummoner.Rank
			summoner.LeaguePoints = updatedSummoner.LeaguePoints
			summoner.Wins = updatedSummoner.Wins
			summoner.Losses = updatedSummoner.Losses
			summoner.IconID = updatedSummoner.IconID
			summoner.Level = updatedSummoner.Level

			updatedCount++
		}

		if err := eu.repository.Save(group); err != nil {
			log.Printf("Error saving group %s: %v", group.Name, err)
			errorCount++
		}
	}

	log.Printf("Elo update completed: %d summoners updated, %d errors", updatedCount, errorCount)
	return nil
}
