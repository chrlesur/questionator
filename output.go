package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

func formatOutput(questions []Question, format string) (string, error) {
	var output strings.Builder

	switch format {
	case "text":
		for i, q := range questions {
			fmt.Fprintf(&output, "Question %d: %s\n", i+1, q.Question)
			fmt.Fprintf(&output, "Réponse: %s\n", q.Answer)
			fmt.Fprintf(&output, "Explication: %s\n", q.Reason)
			fmt.Fprintf(&output, "Passage: %s\n\n", q.Passage)
		}
	case "html":
		funcMap := template.FuncMap{
			"add": func(a, b int) int {
				return a + b
			},
		}
		
		tmpl, err := template.New("html").Funcs(funcMap).Parse(`
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Questions générées par Questionator</title>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 800px; margin: 0 auto; padding: 20px; }
        h1 { color: #2c3e50; }
        .question { background-color: #f9f9f9; border-left: 5px solid #3498db; padding: 15px; margin-bottom: 20px; }
        .question h2 { color: #3498db; margin-top: 0; }
        .answer, .reason, .passage { margin-top: 10px; }
        .answer strong, .reason strong, .passage strong { color: #2c3e50; }
    </style>
</head>
<body>
    <h1>Questions générées par Questionator</h1>
    {{range $index, $question := .}}
    <div class="question">
        <h2>Question {{add $index 1}}</h2>
        <p>{{$question.Question}}</p>
        <div class="answer"><strong>Réponse:</strong> {{$question.Answer}}</div>
        <div class="reason"><strong>Explication:</strong> {{$question.Reason}}</div>
        <div class="passage"><strong>Passage:</strong> {{$question.Passage}}</div>
    </div>
    {{end}}
</body>
</html>
`)
		if err != nil {
			return "", err
		}

		err = tmpl.Execute(&output, questions)
		if err != nil {
			return "", err
		}
	case "markdown":
		for i, q := range questions {
			fmt.Fprintf(&output, "## Question %d\n\n", i+1)
			fmt.Fprintf(&output, "%s\n\n", q.Question)
			fmt.Fprintf(&output, "**Réponse:** %s\n\n", q.Answer)
			fmt.Fprintf(&output, "**Explication:** %s\n\n", q.Reason)
			fmt.Fprintf(&output, "**Passage:** %s\n\n", q.Passage)
		}
	default:
		return "", fmt.Errorf("format de sortie non pris en charge: %s", format)
	}

	return output.String(), nil
}

func writeOutputFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
