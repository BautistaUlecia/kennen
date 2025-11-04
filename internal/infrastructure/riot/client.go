package riot

import (
	"encoding/json"
	"fmt"
	"kennen/internal/domain"
	"net/http"
	"net/url"
)

const accountByRiotIdPath = "/riot/account/v1/accounts/by-riot-id/%s/%s"
const entriesByEncryptedPuuid = "/lol/league/v4/entries/by-puuid/%s"

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    "https://%s.api.riotgames.com",
	}
}
func (c *Client) FindSummoner(gameName, tagLine, region string) (*domain.Summoner, error) {
	acc, err := c.accountByRiotID(gameName, tagLine)
	if err != nil {
		return nil, fmt.Errorf("account error: %w", err)
	}

	entries, err := c.entriesByEncryptedPuuid(acc.Puuid, region)
	if err != nil {
		return nil, fmt.Errorf("entries error: %w", err)
	}

	e := c.findRankedSolo(entries)

	s, err := c.mapToSummoner(e, gameName, tagLine)
	if err != nil {
		return nil, fmt.Errorf("error mapping to summoner: %w", err)
	}
	return s, nil

}

// todo bauti: Deduplicar esto!
func (c *Client) accountByRiotID(gameName string, tagLine string) (AccountDTO, error) {
	acc := AccountDTO{}
	region := "AMERICAS" // Hardcode por ahora, despues mappear LA2, LA1, NA, etc => america, las demas a europe/asia
	gn := url.PathEscape(gameName)
	tl := url.PathEscape(tagLine)

	u := fmt.Sprintf(c.baseURL, region) + fmt.Sprintf(accountByRiotIdPath, gn, tl)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return acc, fmt.Errorf("error generating request for account by riot id %w", err)
	}
	req.Header.Set("X-RIOT-TOKEN", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return acc, fmt.Errorf("error in do request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return acc, fmt.Errorf("riot api responded with code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return acc, fmt.Errorf("error in decode body: %w", err)
	}
	return acc, nil
}
func (c *Client) entriesByEncryptedPuuid(encryptedPuuid, region string) ([]LeagueEntryDTO, error) {
	var entries []LeagueEntryDTO
	p := url.PathEscape(encryptedPuuid)
	u := fmt.Sprintf(c.baseURL, region) + fmt.Sprintf(entriesByEncryptedPuuid, p)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating request for entry by encrypted puuid %w", err)
	}
	req.Header.Set("X-RIOT-TOKEN", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("riot api responded with code %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, fmt.Errorf("error in decode body: %w", err)
	}
	return entries, nil
}

func (c *Client) findRankedSolo(entries []LeagueEntryDTO) LeagueEntryDTO {
	const rankedSolo5x5 = "RANKED_SOLO_5X5"

	for i := range entries {
		if entries[i].QueueType == rankedSolo5x5 {
			return entries[i]
		}
	}
	return entries[0]
}

func (c *Client) mapToSummoner(dto LeagueEntryDTO, gameName, tag string) (*domain.Summoner, error) {
	formattedName := fmt.Sprintf(gameName, tag)

	return &domain.Summoner{
			Name:         formattedName,
			Rank:         dto.Rank,
			LeaguePoints: dto.LeaguePoints,
			Wins:         dto.Wins,
			Losses:       dto.Losses},
		nil
}
