package main

import (
	"log"
)

const (
	version     = "1.0.4"
	programName = "Questionator"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	config := parseFlags()

	log.Println("Lecture du fichier d'entrée...")
	verbatim, err := readInputFile(config.InputFile)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier d'entrée: %v", err)
	}
	inputTokens := countTokens(verbatim)
	log.Printf("Fichier d'entrée lu avec succès. Taille : %d tokens", inputTokens)

	log.Println("Vérification de la clé API...")
	err = verifyAPIKey()
	if err != nil {
		log.Fatalf("Erreur lors de la vérification de la clé API: %v", err)
	}
	log.Println("Clé API vérifiée avec succès")

	log.Printf("Génération de %d questions...", config.QuestionCount)
	questions, err := generateQuestions(verbatim, config.QuestionCount)
	if err != nil {
		log.Fatalf("Erreur lors de la génération des questions: %v", err)
	}
	log.Printf("%d questions générées avec succès", len(questions))

	log.Println("Formatage de la sortie...")
	output, err := formatOutput(questions, config.Format)
	if err != nil {
		log.Fatalf("Erreur lors du formatage de la sortie: %v", err)
	}
	outputTokens := countTokens(output)
	log.Printf("Sortie formatée avec succès. Taille : %d tokens", outputTokens)

	log.Println("Écriture du fichier de sortie...")
	err = writeOutputFile(config.OutputFile, output)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture du fichier de sortie: %v", err)
	}

	log.Printf("Fichier de sortie généré avec succès: %s", config.OutputFile)
}
