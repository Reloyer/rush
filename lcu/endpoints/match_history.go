package endpoints

func GetCurrentSummonerMatchHistory(url string) string {
	return url + "/lol-match-history/v1/products/lol/current-summoner/matches"
}

func GetGame(url string) string {
	return url + "/lol-match-history/v1/games/"
}
