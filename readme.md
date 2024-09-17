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

## Installation de Go

### Windows

1. Téléchargez l'installateur Go pour Windows depuis [le site officiel de Go](https://golang.org/dl/).
2. Exécutez l'installateur et suivez les instructions à l'écran.
3. Ouvrez une nouvelle invite de commande et vérifiez l'installation en tapant :
   ```
   go version
   ```

### Linux

1. Ouvrez un terminal.
2. Téléchargez et installez Go en utilisant les commandes suivantes (remplacez `1.16.5` par la version la plus récente) :
   ```
   wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
   ```
3. Ajoutez Go à votre PATH en ajoutant ces lignes à votre `~/.bashrc` ou `~/.zshrc` :
   ```
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```
4. Rechargez votre profil de shell :
   ```
   source ~/.bashrc   # ou source ~/.zshrc si vous utilisez zsh
   ```
5. Vérifiez l'installation :
   ```
   go version
   ```

### macOS

1. Téléchargez le package d'installation pour macOS depuis [le site officiel de Go](https://golang.org/dl/).
2. Ouvrez le package téléchargé et suivez les instructions d'installation.
3. Ouvrez un nouveau terminal et vérifiez l'installation :
   ```
   go version
   ```

Alternativement, si vous utilisez Homebrew, vous pouvez installer Go avec :
```
brew install go
```

## Installation de Questionator

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

Ce projet est sous licence GNU General Public License v3.0 (GPL-3.0). Cela signifie que vous êtes libre de :

- Utiliser ce logiciel pour n'importe quel usage
- Modifier ce logiciel
- Distribuer ce logiciel

Cependant, si vous distribuez une version modifiée de ce logiciel, vous devez :

- Publier votre code source
- Licencier votre version sous la GPL-3.0
- Documenter les changements que vous avez effectués

Pour plus de détails, consultez le fichier `LICENSE` inclus dans ce projet ou visitez [https://www.gnu.org/licenses/gpl-3.0.en.html](https://www.gnu.org/licenses/gpl-3.0.en.html).
