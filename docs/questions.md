# Module questions

Le module `questions` gère la génération et le traitement des questions à partir du texte d'entrée.

## Fonctions principales

### `generateQuestions(verbatim string, questionCount int) ([]Question, error)`

Génère un nombre spécifié de questions basées sur le texte d'entrée en utilisant l'API Claude.

Paramètres :
- `verbatim` : Le texte d'entrée
- `questionCount` : Le nombre de questions à générer

Retourne un tableau de `Question` et une erreur éventuelle.

### `parseTextResponse(text string) ([]Question, error)`

Parse la réponse textuelle de l'API Claude pour extraire les questions, réponses, explications et passages associés.

Paramètres :
- `text` : La réponse textuelle de l'API

Retourne un tableau de `Question` et une erreur éventuelle.

## Structures

### `Question`

Représente une question générée avec sa réponse, son explication et le passage associé.

Champs :
- `Question` : La question
- `Answer` : La réponse
- `Reason` : L'explication
- `Passage` : Le passage associé du texte original

## Utilisation

```go
questions, err := generateQuestions(verbatim, 10)
if err != nil {
    log.Fatalf("Erreur lors de la génération des questions: %v", err)
}
```

Ce module est crucial pour la fonctionnalité principale de Questionator, transformant le texte d'entrée en un ensemble de questions structurées.
