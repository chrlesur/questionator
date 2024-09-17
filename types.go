package main

const (
	ClaudeModelName = "claude-3-5-sonnet-20240620"
)

// Le reste du fichier reste inchangé
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Reason   string `json:"reason"`
	Passage  string `json:"passage"`
}

type APIRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}
