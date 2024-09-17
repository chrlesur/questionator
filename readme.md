# Questionator

Questionator est un outil en ligne de commande qui génère automatiquement des questions à partir d'un texte donné en utilisant l'API Claude d'Anthropic. Il est conçu pour créer des questionnaires pertinents basés sur le contenu fourni, ce qui peut être utile pour des fins éducatives, de révision ou de test de compréhension.

## Fonctionnalités

- Génération de questions basées sur un texte d'entrée
- Utilisation de l'API Claude 3.5 Sonnet d'Anthropic pour une génération de questions de haute qualité
- Personnalisation du nombre de questions générées
- Support de différents formats de sortie (texte, HTML, Markdown)

## Prérequis

- Go 1.16 ou version ultérieure
- Une clé API valide pour l'API Claude d'Anthropic

## Installation

1. Clonez ce dépôt :
   ```
   git clone https://github.com/votre-nom/questionator.git
   cd questionator
   ```

2. Compilez le programme :
   ```
   go build -o questionator
   ```

## Configuration

Avant d'utiliser Questionator, vous devez définir votre clé API Anthropic. Vous pouvez le faire en définissant une variable d'environnement :

- Sur Windows (PowerShell) :
  ```
  $env:ANTHROPIC_API_KEY="votre_clé_api_ici"
  ```
- Sur Linux/macOS :
  ```
  export ANTHROPIC_API_KEY="votre_clé_api_ici"
  ```

## Utilisation

Utilisez la commande suivante pour générer des questions :

```
.\questionator.exe -input chemin/vers/votre/fichier.txt -output questions.html -format html -questions 10
```

Options disponibles :
- `-input` : Chemin vers le fichier d'entrée contenant le texte à analyser (obligatoire)
- `-output` : Nom du fichier de sortie (par défaut : "output.html")
- `-format` : Format de sortie (options : "text", "html", "markdown", par défaut : "html")
- `-questions` : Nombre de questions à générer (par défaut : 10, maximum : 15)
- `-version` : Affiche la version du programme

## Structure du projet

- `main.go` : Point d'entrée du programme
- `config.go` : Gestion de la configuration et des drapeaux
- `api.go` : Interactions avec l'API Claude
- `questions.go` : Génération et traitement des questions
- `output.go` : Formatage et écriture de la sortie
- `utils.go` : Fonctions utilitaires
- `types.go` : Définitions des types de données

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou à soumettre une pull request.

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.
