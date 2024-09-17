# Module output

Le module `output` gère le formatage et l'écriture de la sortie générée par Questionator.

## Fonctions principales

### `formatOutput(questions []Question, format string) (string, error)`

Formate les questions générées selon le format spécifié.

Paramètres :
- `questions` : Tableau de questions à formater
- `format` : Format de sortie ("text", "html", ou "markdown")

Retourne une chaîne formatée et une erreur éventuelle.

### `writeOutputFile(filename string, content string) error`

Écrit le contenu formaté dans un fichier.

Paramètres :
- `filename` : Nom du fichier de sortie
- `content` : Contenu à écrire

Retourne une erreur éventuelle.

## Formats supportés

- Texte : Format simple avec chaque question et ses détails sur des lignes séparées
- HTML : Format structuré avec des styles CSS pour une meilleure lisibilité
- Markdown : Format compatible avec les systèmes de rendu Markdown

## Utilisation

```go
output, err := formatOutput(questions, config.Format)
if err != nil {
    log.Fatalf("Erreur lors du formatage de la sortie: %v", err)
}

err = writeOutputFile(config.OutputFile, output)
if err != nil {
    log.Fatalf("Erreur lors de l'écriture du fichier de sortie: %v", err)
}
```

Ce module est responsable de la présentation finale des questions générées et de leur sauvegarde dans un fichier.
