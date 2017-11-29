package address

import "regexp"

type Module struct {
	special, ilegalOctalNumber, newLine *regexp.Regexp
}
