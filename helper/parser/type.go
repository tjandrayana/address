package parser

import "regexp"

type Parser struct {
	punc, decimal, zero, unicode, integer, special, ilegalOctalNumber, newLine, alphabeth, number, email, number_only *regexp.Regexp
}
