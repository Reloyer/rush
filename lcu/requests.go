package lcu

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Reloyer/rush/lcu/api/endpoints"
	"github.com/Reloyer/rush/lcu/api/types"
)

type GetReq struct {
	URL            string
	AuthToken      string
	AdditionalArgs map[string]interface{}
}

func getRequest(g GetReq) (*http.Response, error) {
	for key, value := range g.AdditionalArgs {
		g.URL += fmt.Sprintf("&%s=%v", key, value)
	}
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

func GetCurrentSummoner(url string, authToken string) (summ types.Summoner, err error) {
	url = endpoints.GetCurrentSummoner(url)
	g := GetReq{
		URL:       url,
		AuthToken: authToken,
	}
	response, err := getRequest(g)
	if err != nil {
		return types.Summoner{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return types.Summoner{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&summ)
	if err != nil {
		return types.Summoner{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return summ, nil
}

func GetCurrentRankedStats(url string, authToken string) (rank types.Ranked, err error) {
	url = endpoints.GetCurrentRankedStats(url)
	g := GetReq{
		URL:       url,
		AuthToken: authToken,
	}
	response, err := getRequest(g)
	if err != nil {
		return types.Ranked{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return types.Ranked{}, fmt.Errorf("request failed with status: %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&rank)
	if err != nil {
		return types.Ranked{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return rank, nil
}
func GetCurrentSummonerMatches(url string, authToken string, puuid int, begIndex int, endIndex int) (matches types.Matches, err error) {
	url = endpoints.GetCurrentSummonerMatches(url)
	req := GetReq{
		URL:       url,
		AuthToken: authToken,
		AdditionalArgs: map[string]interface{}{
			"PUUID":    puuid,
			"BegIndex": begIndex,
			"EndIndex": endIndex,
		},
	}

	response, err := getRequest(req)
	if err != nil {
		return types.Matches{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return types.Matches{}, fmt.Errorf("request failed with status: %s", response.Status)
	}
	err = json.NewDecoder(response.Body).Decode(&matches)

	if err != nil {
		return types.Matches{}, fmt.Errorf("error decoding JSON: %s", err)
	}

	return matches, nil
}
