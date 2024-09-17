# Module config

Le module `config` gère la configuration et le parsing des flags de ligne de commande pour Questionator.

## Structures

### `Config`

Contient la configuration du programme :
- `InputFile` : Chemin du fichier d'entrée
- `OutputFile` : Chemin du fichier de sortie
- `Format` : Format de sortie (text, html, markdown)
- `QuestionCount` : Nombre de questions à générer

## Fonctions

### `parseFlags() Config`

Parse les flags de ligne de commande et retourne une instance de `Config`.

Flags supportés :
- `-input` : Fichier d'entrée (obligatoire)
- `-output` : Fichier de sortie (par défaut : "output.html")
- `-format` : Format de sortie (par défaut : "html")
- `-questions` : Nombre de questions (par défaut : 10, max 15)
- `-version` : Affiche la version du programme

## Utilisation

```go
config := parseFlags()
```

Cette fonction est généralement appelée au début de l'exécution du programme pour configurer les paramètres d'exécution.
