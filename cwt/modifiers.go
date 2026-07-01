package cwt

import (
	"strings"

	"bahmut.de/pdx-documentation-manager/parser"
)

func PrintModifiers(docs *parser.ScriptDocumentation) string {
	var builder = strings.Builder{}

	builder.WriteString("# Auto Generated:\n")
	builder.WriteString("# This is a list of all base game modifiers.\n")
	builder.WriteString("modifiers = {\n")

	for _, modifier := range docs.ModifierDocumentation.Elements {
		builder.WriteString("\t")
		builder.WriteString(modifier.Name)
		builder.WriteString(" = ")
		builder.WriteString(modifier.Mask)
		builder.WriteString("\n")
	}

	builder.WriteString("}")

	return builder.String()
}

func PrintModifierCategories(docs *parser.ScriptDocumentation) string {
	var builder = strings.Builder{}

	builder.WriteString("# Auto Generated:\n")
	builder.WriteString("# This is a list of all base game modifier masks.\n")
	builder.WriteString("modifier_categories = {\n")

	var masks = make(map[string]string)

	for _, modifier := range docs.ModifierDocumentation.Elements {
		if masks[modifier.Mask] != "" {
			continue
		}
		builder.WriteString("\t")
		builder.WriteString(modifier.Mask)
		builder.WriteString(" = {\n")
		builder.WriteString("\t\tsupported_scopes = { any }\n")
		builder.WriteString("\t}\n")
		masks[modifier.Mask] = modifier.Mask
	}

	builder.WriteString("}")

	return builder.String()
}
