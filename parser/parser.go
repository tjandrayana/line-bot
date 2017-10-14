package parser

import "regexp"

func New() *Parser {
	parser := &Parser{
		punc:              regexp.MustCompile(`\W`),
		decimal:           regexp.MustCompile(`\D`),
		zero:              regexp.MustCompile(`^0`),
		unicode:           regexp.MustCompile(`[\pL\p{Mc}\p{Mn}-_']+`),
		integer:           regexp.MustCompile(`^(?:[-+]?(?:0|[1-9][0-9]*))$`),
		Special:           regexp.MustCompile(`[_+-.:,!@#$%^&*();\\/|<>"'?=]`),
		ilegalOctalNumber: regexp.MustCompile(` 0+`),
		newLine:           regexp.MustCompile("\\[.*?]"),
		alphabeth:         regexp.MustCompile("[^A-Za-z]+"),
		number:            regexp.MustCompile("[0-9]+"),
	}

	return parser
}
