# Module main

Le module `main` est le point d'entrée du programme Questionator. Il orchestre le flux d'exécution principal de l'application.

## Fonctions principales

### `main()`

Cette fonction est le point d'entrée du programme. Elle effectue les opérations suivantes :

1. Configure la journalisation
2. Parse les flags de ligne de commande
3. Lit le fichier d'entrée
4. Vérifie la clé API
5. Génère les questions
6. Formate la sortie
7. Écrit le fichier de sortie

## Dépendances

- `config` : Pour le parsing des flags
- `api` : Pour la vérification de la clé API
- `questions` : Pour la génération des questions
- `output` : Pour le formatage et l'écriture de la sortie
- `utils` : Pour diverses fonctions utilitaires

## Constantes

- `version` : Version actuelle du programme
- `programName` : Nom du programme

## Utilisation

Ce module est automatiquement exécuté lorsque le programme est lancé. Il ne nécessite pas d'être importé ou appelé directement par d'autres parties du code.
