package parser

import (
	"regexp"
)

func New() *Parser {
	parser := &Parser{
		special:           regexp.MustCompile(`[_+-.:,!@#$%^&*();\\/|<>"'?=]`),
		ilegalOctalNumber: regexp.MustCompile(` 0+`),
		newLine:           regexp.MustCompile("\\[.*?]"),
	}
	province = make(map[string]string)
	city = make(map[string]string)
	district = make(map[string]string)

	parser.LoadStopWords()

	return parser
}
