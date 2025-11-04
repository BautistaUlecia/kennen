package riot

// /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}

// /lol/league/v4/entries/by-puuid/{encryptedPUUID

type AccountDTO struct {
	Puuid string
}
type LeagueEntryDTO struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}
