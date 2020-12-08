package parser

import (
	"fmt"

	"github.com/alecthomas/participle"
)

var parser = participle.MustBuild(&LuggageRule{})

type LuggageRule struct {
	Color    *Color     `@@ "bags" "contain" `
	EdgeList *EdgeList  `@@?`
	Empty    *EmptyRule `@@?`
}

type EdgeList struct {
	Edges *[]Edge `@@*`
}

type EmptyRule struct {
	Empty bool `@"no" "other" "bags""."?`
}

type Edge struct {
	Count int    `@Int`
	Color *Color `@@ ("bag" | "bags")","?"."?" "?`
}

type Color struct {
	Adjective string `@Ident`
	Name      string `@Ident`
}

func ParseLine(line string) (LuggageRule, error) {
	rule := LuggageRule{}

	err := parser.ParseString(line, &rule)

	if err != nil {
		return LuggageRule{}, fmt.Errorf("Error parsing: %v", err.Error())
	}

	return rule, nil
}
