package cwt

import (
	"strings"

	"bahmut.de/pdx-documentation-manager/parser"
)

func PrintScopeEnum(docs *parser.ScriptDocumentation) string {
	var builder = strings.Builder{}

	builder.WriteString("# Auto Generated:\n")
	builder.WriteString("# This is an enum of all scopes.\n")
	builder.WriteString("enum = {\n")
	builder.WriteString("\tenum[scopes] = {\n")

	for _, scope := range docs.ScopeDocumentation.Elements {
		builder.WriteString("\t\t")
		builder.WriteString(scope.Name)
		builder.WriteString("\n")
	}

	builder.WriteString("\t}\n}")

	return builder.String()
}

func PrintScopes(docs *parser.ScriptDocumentation) string {
	var builder = strings.Builder{}

	builder.WriteString("# Auto Generated:\n")
	builder.WriteString("# This is a list of all scopes.\n")
	builder.WriteString("scopes = {\n")

	for _, scope := range docs.ScopeDocumentation.Elements {
		builder.WriteString("\t")
		builder.WriteString(scope.Name)
		builder.WriteString(" = {\n")
		builder.WriteString("\t\taliases = { ")
		builder.WriteString(scope.Name)
		builder.WriteString(" }\n")
		builder.WriteString("\t}\n")
	}

	builder.WriteString("}")

	return builder.String()
}
