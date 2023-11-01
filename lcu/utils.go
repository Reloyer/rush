package lcu

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type GetReq struct {
	url            string
	authtoken      string
	AdditionalArgs map[string]interface{}
}

func Get(url, authToken string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Authorization", "Basic "+authToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %s", err)
	}

	return response, nil
}
