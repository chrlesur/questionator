package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func generateQuestions(verbatim string, questionCount int) ([]Question, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("la clé API ANTHROPIC_API_KEY n'est pas définie")
	}

	prompt := fmt.Sprintf("Générez %d questions pertinentes basées sur l'ensemble du verbatim suivant. Pour chaque question, fournissez la bonne réponse, une explication de pourquoi c'est la bonne réponse, et le passage associé du verbatim. Verbatim: %s", questionCount, verbatim)

	reqBody, err := json.Marshal(APIRequest{
		Model:     ClaudeModelName,
		MaxTokens: 8000,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la requête JSON: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la requête HTTP: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'envoi de la requête à l'API: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse de l'API: %v", err)
	}

	log.Printf("Réponse reçue. Taille : %d tokens", countTokens(string(body)))
	log.Printf("Contenu de la réponse : %s", string(body))

	var apiResp APIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage de la réponse JSON de l'API: %v", err)
	}

	if len(apiResp.Content) == 0 {
		return nil, fmt.Errorf("la réponse de l'API ne contient pas de contenu")
	}

	questions, err := parseTextResponse(apiResp.Content[0].Text)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'extraction des questions: %v", err)
	}

	// S'assurer que nous n'avons pas plus de questions que demandé
	if len(questions) > questionCount {
		questions = questions[:questionCount]
	}

	return questions, nil
}

func parseTextResponse(text string) ([]Question, error) {
	questions := []Question{}
	lines := strings.Split(text, "\n")

	currentQuestion := Question{}
	currentField := ""
	questionPattern := regexp.MustCompile(`^\d+\.\s`)
	qrPattern := regexp.MustCompile(`Q:\s*(.*?)\s*R:\s*(.*)`)
	passagePattern := regexp.MustCompile(`Passage(?:\s*associé)?\s*:\s*"?(.+?)"?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if questionPattern.MatchString(line) {
			if currentQuestion.Question != "" {
				questions = append(questions, currentQuestion)
			}
			currentQuestion = Question{}
			currentField = "question"

			if match := qrPattern.FindStringSubmatch(line); match != nil {
				currentQuestion.Question = strings.TrimSpace(match[1])
				currentQuestion.Answer = strings.TrimSpace(match[2])
				currentField = "answer"
			} else {
				currentQuestion.Question = strings.TrimSpace(questionPattern.ReplaceAllString(line, ""))
			}
		} else if strings.HasPrefix(line, "Question :") {
			currentQuestion.Question = strings.TrimPrefix(line, "Question :")
			currentField = "question"
		} else if strings.HasPrefix(line, "Réponse:") || strings.HasPrefix(line, "Réponse :") {
			currentQuestion.Answer = strings.TrimPrefix(strings.TrimPrefix(line, "Réponse:"), "Réponse :")
			currentField = "answer"
		} else if strings.HasPrefix(line, "Explication:") || strings.HasPrefix(line, "Explication :") {
			currentQuestion.Reason = strings.TrimPrefix(strings.TrimPrefix(line, "Explication:"), "Explication :")
			currentField = "reason"
		} else if strings.HasPrefix(line, "Passage:") || strings.HasPrefix(line, "Passage :") {
			currentQuestion.Passage = strings.TrimPrefix(strings.TrimPrefix(line, "Passage:"), "Passage :")
			currentField = "passage"
		} else {
			switch currentField {
			case "question":
				currentQuestion.Question += " " + line
			case "answer":
				currentQuestion.Answer += " " + line
			case "reason":
				if match := passagePattern.FindStringSubmatch(line); match != nil {
					currentQuestion.Passage = match[1]
					currentField = "passage"
				} else {
					currentQuestion.Reason += " " + line
				}
			case "passage":
				currentQuestion.Passage += " " + line
			}
		}
	}

	// Ajouter la dernière question
	if currentQuestion.Question != "" {
		questions = append(questions, currentQuestion)
	}

	// Nettoyage final
	for i := range questions {
		questions[i].Question = strings.TrimSpace(questions[i].Question)
		questions[i].Answer = strings.TrimSpace(questions[i].Answer)
		questions[i].Reason = strings.TrimSpace(questions[i].Reason)
		questions[i].Passage = strings.TrimSpace(questions[i].Passage)
	}

	if len(questions) == 0 {
		return nil, fmt.Errorf("aucune question n'a pu être extraite de la réponse")
	}

	return questions, nil
}
