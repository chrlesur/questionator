package main

import (
	"io/ioutil"
	"strings"
	"unicode"
)

func readInputFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func countTokens(text string) int {
	// Règles de tokenisation approximatives
	tokens := 0
	inToken := false
	for _, r := range text {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			if !inToken {
				tokens++
				inToken = true
			}
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			// La plupart des ponctuations sont des tokens séparés
			tokens++
			inToken = false
		case unicode.IsSpace(r):
			inToken = false
		}
	}

	// Ajustement pour les tokens spéciaux et les sous-mots
	words := strings.Fields(text)
	for _, word := range words {
		// Gestion des contractions courantes en anglais
		if strings.Contains(word, "'") {
			tokens++ // Ajoute un token pour la contraction
		}
		// Gestion des mots composés
		if strings.Contains(word, "-") {
			tokens += strings.Count(word, "-")
		}
		// Gestion des préfixes et suffixes courants
		prefixes := []string{"un", "re", "dis"}
		suffixes := []string{"ed", "ing", "ly", "s", "es"}
		for _, prefix := range prefixes {
			if strings.HasPrefix(strings.ToLower(word), prefix) && len(word) > len(prefix) {
				tokens++
			}
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(strings.ToLower(word), suffix) && len(word) > len(suffix) {
				tokens++
			}
		}
	}

	return tokens
}
