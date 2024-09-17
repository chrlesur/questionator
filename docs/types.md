# Module types

Le module `types` définit les structures de données et les constantes utilisées dans l'ensemble du projet Questionator.

## Constantes

- `ClaudeModelName` : Nom du modèle Claude utilisé pour la génération de questions

## Structures

### `Question`

Représente une question générée avec sa réponse, son explication et le passage associé.

Champs :
- `Question` : string
- `Answer` : string
- `Reason` : string
- `Passage` : string

### `APIRequest`

Représente une requête à l'API Claude.

Champs :
- `Model` : string
- `MaxTokens` : int
- `Messages` : []Message

### `Message`

Représente un message dans la conversation avec Claude.

Champs :
- `Role` : string
- `Content` : string

### `APIResponse`

Représente la réponse de l'API Claude.

Champs :
- `Content` : []struct { Text string }

## Utilisation

Ces types sont utilisés dans l'ensemble du projet pour structurer les données et interagir avec l'API Claude.

```go
var question Question
var apiReq APIRequest
var apiResp APIResponse
```

Ce module centralise les définitions de types, facilitant la maintenance et assurant la cohérence dans tout le projet.
