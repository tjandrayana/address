package parser

import "regexp"

type Parser struct {
	special, ilegalOctalNumber, newLine *regexp.Regexp
}
