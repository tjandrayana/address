package address

import (
	"regexp"
)

func New() *Module {
	m := &Module{
		special:           regexp.MustCompile(`[_+-.:,!@#$%^&*();\\/|<>"'?=]`),
		ilegalOctalNumber: regexp.MustCompile(` 0+`),
		newLine:           regexp.MustCompile("\\[.*?]"),
	}
	province = make(map[string]string)
	city = make(map[string]string)
	district = make(map[string]string)

	m.LoadStopWords()

	return m
}
