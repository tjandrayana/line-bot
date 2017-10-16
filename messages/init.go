package messages

import (
	"github.com/robfig/cron"
	"github.com/tjandrayana/line-bot/parser"
)

var p *parser.Parser

type module struct {
	cron *cron.Cron
}

var m module

func Init() {

	p = parser.New()

	// m.cron = cron.New()
	// m.cron.AddFunc("@every 1m", func() {
	// 	m.doJob()
	// })

	// m.cron.Start()

}
