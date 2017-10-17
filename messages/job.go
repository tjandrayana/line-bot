package messages

func (m module) doJob() error {
	dat := Data{}
	msg := Message{
		Type: "text",
		Text: Horoskop(),
	}
	arrMsg := []Message{msg}
	SendMessages(dat, arrMsg)

	return nil
}
