package cwt

import (
	"strings"

	"bahmut.de/pdx-documentation-manager/parser"
)

var eventScopes = map[string]string{
	"country":   "country",
	"state":     "state",
	"character": "character",
}

func PrintOnActions(docs *parser.ScriptDocumentation) string {
	var builder = strings.Builder{}

	builder.WriteString("# Auto Generated:\n")
	builder.WriteString("# This is a list of all base game on actions.\n")
	builder.WriteString("on_actions = {\n")

	for _, onAction := range docs.OnActionDocumentation.Elements {
		if !onAction.FromCode {
			continue
		}
		builder.WriteString("\t## replace_scopes = { this = ")
		builder.WriteString(onAction.Scope)
		builder.WriteString(" root = ")
		builder.WriteString(onAction.Scope)
		builder.WriteString(" }\n")
		builder.WriteString("\t## event_type = ")
		if eventScopes[onAction.Scope] != "" {
			builder.WriteString(eventScopes[onAction.Scope])
		} else {
			builder.WriteString("scopeless")
		}
		builder.WriteString("\n\t")
		builder.WriteString(onAction.Name)

		builder.WriteString("\n\n")
	}

	builder.WriteString("}")

	return builder.String()
}
