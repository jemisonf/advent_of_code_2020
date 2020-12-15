package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

var lexer = stateful.MustSimple([]stateful.Rule{
	{"Mask", "[01X]{36}", nil},
	{"Number", `[0-9]+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	{"EOL", `[\n\r]+`, nil},
	{"whitespace", `[ \t\n]+`, nil},
})
var parser = participle.MustBuild(&Program{}, participle.Lexer(lexer))

type Program struct {
	Steps []Step `@@*`
}

type Step struct {
	Bitmask     Bitmask       `@@`
	Instruction []Instruction `@@*`
}

type Bitmask struct {
	Mask string `"mask" "=" @Mask "\n"`
}

type Instruction struct {
	Address int `"mem" "[" @Number "]"`
	Value   int `"=" @Number "\n"?`
}

func ParseProgram(program string) (Program, error) {
	p := Program{}

	if err := parser.ParseString("", program, &p); err != nil {
		return Program{}, err
	}

	return p, nil
}
