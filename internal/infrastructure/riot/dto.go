package riot

// /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}

// /lol/league/v4/entries/by-puuid/{encryptedPUUID

type AccountDTO struct {
	Puuid string
}
type LeagueEntryDTO struct {
	Rank         string
	Leaguepoints int
	wins         int
	loses        int
}
