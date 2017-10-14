package messages

type Data struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64 `json:"timestamp"`
		Message   struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Reply struct {
	ReplyToken string    `json:"replyToken,omitempty"`
	Messages   []Message `json:"messages"`
}

type PushMessage struct {
	Messages []Message `json:"messages"`
	To       string    `json:"to"`
}
