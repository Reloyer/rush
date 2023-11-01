package lcu

import (
	"encoding/json"
	"fmt"
	"github.com/Reloyer/rush/lcu/api/endpoints"
	"github.com/Reloyer/rush/lcu/api/types"
	"net/http"
)

func GetCurrentSummoner(url string, authToken string) (summ types.Summoner, err error) {
	url = endpoints.GetCurrentSummoner(url)

	response, err := Get(url, authToken)
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

	response, err := Get(url, authToken)
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
