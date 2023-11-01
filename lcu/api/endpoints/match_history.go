package endpoints

func GetCurrentSummonerMatches(url string) string {
	return url + "/lol-match-history/v1/products/lol/current-summoner/matches"
}
