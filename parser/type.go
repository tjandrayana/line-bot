package parser

import "regexp"

type Parser struct {
	punc, decimal, zero, unicode, integer, Special, ilegalOctalNumber, newLine, alphabeth, number *regexp.Regexp
}
