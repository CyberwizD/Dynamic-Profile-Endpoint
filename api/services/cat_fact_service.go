package services

import (
	"encoding/json"
	"net/http"
	"time"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func GetCatFact() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://catfact.ninja/fact")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var catFact CatFact
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		return "", err
	}

	return catFact.Fact, nil
}
