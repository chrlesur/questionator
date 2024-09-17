# Module utils

Le module `utils` contient diverses fonctions utilitaires utilisées dans Questionator.

## Fonctions

### `readInputFile(filename string) (string, error)`

Lit le contenu d'un fichier et le retourne sous forme de chaîne.

Paramètres :
- `filename` : Chemin du fichier à lire

Retourne le contenu du fichier et une erreur éventuelle.

### `countTokens(text string) int`

Estime le nombre de tokens dans un texte donné.

Paramètres :
- `text` : Texte à analyser

Retourne une estimation du nombre de tokens.

## Utilisation

```go
content, err := readInputFile(config.InputFile)
if err != nil {
    log.Fatalf("Erreur lors de la lecture du fichier d'entrée: %v", err)
}

tokenCount := countTokens(content)
log.Printf("Nombre de tokens dans le fichier d'entrée : %d", tokenCount)
```

Ce module fournit des fonctionnalités de support utilisées dans différentes parties de Questionator.
