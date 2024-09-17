# Module api

Le module `api` gère les interactions avec l'API Claude d'Anthropic.

## Constantes

- `apiURL` : URL de l'API Claude

## Fonctions

### `verifyAPIKey() error`

Vérifie si la clé API est valide en effectuant une requête de test à l'API Claude.

Retourne une erreur si la clé API n'est pas définie ou si la requête de test échoue.

## Utilisation

```go
err := verifyAPIKey()
if err != nil {
    log.Fatalf("Erreur lors de la vérification de la clé API: %v", err)
}
```

Cette fonction est généralement appelée au début du programme pour s'assurer que la clé API est valide avant de procéder à la génération des questions.

## Dépendances

Ce module dépend de la variable d'environnement `ANTHROPIC_API_KEY` qui doit être définie avec une clé API Claude valide.
