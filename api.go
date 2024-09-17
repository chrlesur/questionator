package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const apiURL = "https://api.anthropic.com/v1/messages"

func verifyAPIKey() error {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("la clé API ANTHROPIC_API_KEY n'est pas définie")
	}

	reqBody, err := json.Marshal(APIRequest{
		Model:     ClaudeModelName,
		MaxTokens: 100,
		Messages: []Message{
			{Role: "user", Content: "Ceci est un test de vérification de la clé API."},
		},
	})
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête JSON de test: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête HTTP de test: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête de test à l'API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("la requête de test a échoué avec le code de statut %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
