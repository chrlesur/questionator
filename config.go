package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	InputFile      string
	OutputFile     string
	Format         string
	QuestionCount  int
}

func parseFlags() Config {
	inputFile := flag.String("input", "", "Fichier d'entrée contenant le verbatim")
	outputFile := flag.String("output", "output.html", "Fichier de sortie pour les questions générées")
	format := flag.String("format", "html", "Format de sortie (text, html, markdown)")
	questionCount := flag.Int("questions", 10, "Nombre de questions à générer (maximum 15)")
	showVersion := flag.Bool("version", false, "Afficher la version du programme")

	flag.Usage = func() {
		fmt.Printf("Usage de %s:\n", programName)
		fmt.Printf("  %s [options]\n\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("%s version %s\n", programName, version)
		os.Exit(0)
	}

	if *inputFile == "" {
		fmt.Println("Erreur : Le fichier d'entrée est requis.")
		flag.Usage()
		os.Exit(1)
	}

	if *questionCount > 15 {
		fmt.Println("Attention : Le nombre maximum de questions est limité à 15. La valeur a été ajustée.")
		*questionCount = 15
	}

	return Config{
		InputFile:      *inputFile,
		OutputFile:     *outputFile,
		Format:         *format,
		QuestionCount:  *questionCount,
	}
}
