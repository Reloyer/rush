package lcu

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Reloyer/rush/lcu/datatypes"
	"github.com/Reloyer/rush/lcu/endpoints"
)

type GetReq struct {
	URL            string
	AuthToken      string
	AdditionalArgs map[string]interface{}
}

func getRequest(g GetReq) (*http.Response, error) {
	u, err := url.Parse(g.URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %s", err)
	}

	q := u.Query()
	for key, value := range g.AdditionalArgs {
		q.Set(key, fmt.Sprintf("%v", value))
	}
	u.RawQuery = q.Encode()
	g.URL = u.String()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Authorization", "Basic "+g.AuthToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %s", err)
	}

	return response, nil
}

func GetCurrentSummoner(url string, authToken string) (summ datatypes.Summoner, err error) {
	url = endpoints.GetCurrentSummoner(url)
	g := GetReq{
		URL:       url,
		AuthToken: authToken,
	}
	response, err := getRequest(g)
	if err != nil {
		return datatypes.Summoner{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return datatypes.Summoner{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&summ)
	if err != nil {
		return datatypes.Summoner{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return summ, nil
}

func GetCurrentRankedStats(url string, authToken string) (rank datatypes.Ranked, err error) {
	url = endpoints.GetCurrentRankedStats(url)
	g := GetReq{
		URL:       url,
		AuthToken: authToken,
	}
	response, err := getRequest(g)
	if err != nil {
		return datatypes.Ranked{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return datatypes.Ranked{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&rank)
	if err != nil {
		return datatypes.Ranked{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return rank, nil
}

func GetCurrentSummonerMatchHistory(url string, authToken string, begIndex int, endIndex int) (matchHistory datatypes.MatchHistory, err error) {
	url = endpoints.GetCurrentSummonerMatchHistory(url)
	req := GetReq{
		URL:       url,
		AuthToken: authToken,
		AdditionalArgs: map[string]interface{}{
			"begIndex": begIndex,
			"endIndex": endIndex,
		},
	}

	response, err := getRequest(req)
	if err != nil {
		return datatypes.MatchHistory{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return datatypes.MatchHistory{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&matchHistory)
	if err != nil {
		return datatypes.MatchHistory{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return matchHistory, nil
}
func GetGame(url string, authToken string, gameId int) (game datatypes.Game, err error) {
	url = endpoints.GetGame(url) + fmt.Sprintf("%d", gameId)
	req := GetReq{
		URL:       url,
		AuthToken: authToken,
	}
	response, err := getRequest(req)
	if err != nil {
		return datatypes.Game{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println(err)
		return datatypes.Game{}, err
	}

	err = json.NewDecoder(response.Body).Decode(&game)

	if err != nil {
		return datatypes.Game{}, err
	}
	return game, nil
}
