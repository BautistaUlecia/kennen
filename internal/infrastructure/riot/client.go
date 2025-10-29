package riot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}

// /lol/league/v4/entries/by-puuid/{encryptedPUUID

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
		baseURL:    "https://americas.api.riotgames.com",
	}
}
func (c *Client) AccountByRiotID(gameName string, tagLine string) (AccountDTO, error) {
	acc := AccountDTO{}
	gn := url.PathEscape(gameName)
	tl := url.PathEscape(tagLine)
	u := c.baseURL + fmt.Sprintf(accountByRiotIdPath, gn, tl)

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
