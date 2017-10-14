package messages

import "github.com/tjandrayana/line-bot/parser"

var p *parser.Parser

func Init() {

	p = parser.New()

}
