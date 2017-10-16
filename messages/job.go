package messages

import (
	"fmt"
	"time"
)

func (m module) doJob() error {
	dat := Data{}
	msg := Message{
		Type: "text",
		Text: fmt.Sprintf("Sekarang jam %s\n", time.Now().String()),
	}
	arrMsg := []Message{msg}
	SendMessages(dat, arrMsg)

	return nil
}
