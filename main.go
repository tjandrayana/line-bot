package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjandrayana/line-bot/utils"
)

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

// var bot *linebot.Client

func main() {
	// bot, err := linebot.New(os.Getenv("channelSecret"), os.Getenv("channelAccessToken"))

	// log.Println("Bot:", bot, " err:", err)

	r := gin.New()
	r.GET("/ping", ping)

	r.POST("/line/triger", Triger)

	r.Run()
}

func ping(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ParseRequest func
func Triger(c *gin.Context) {
	defer c.Request.Body.Close()
	channelSecret := os.Getenv("channelSecret")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("error : ", err)
		return
	}

	fmt.Printf("\n%+v\n", string(body))

	if !validateSignature(channelSecret, c.Request.Header.Get("X-Line-Signature"), body) {
		log.Println("error : ", err)
		return
	}

	var dat Data
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Println(err)
	}

	if err := reply(dat); err != nil {
		log.Println("Reply ERROR = ", err)
	}

	if err := pushMessage(dat); err != nil {
		log.Println("Push Message ERROR = ", err)
	}

	fmt.Println("\nSuccess\n")

}

func validateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))
	hash.Write(body)
	return hmac.Equal(decoded, hash.Sum(nil))
}

type Reply struct {
	ReplyToken string    `json:"replyToken,omitempty"`
	Messages   []Message `json:"messages"`
}

type PushMessage struct {
	Messages []Message `json:"messages"`
	To       string    `json:"to"`
}

func reply(dat Data) error {

	mess1 := Message{
		Type: "text",
		Text: "Hai User ... ",
	}

	mess2 := Message{
		Type: "text",
		Text: "May I help you ...? ",
	}

	messages := []Message{mess1, mess2}

	rep := Reply{
		ReplyToken: dat.Events[0].ReplyToken,
		Messages:   messages,
	}

	agent := utils.NewHTTPRequest()
	agent.Url = "https://api.line.me"
	agent.Path = "/v2/bot/message/reply"
	agent.Method = "POST"
	agent.IsJson = true
	agent.Json = rep

	agent.Headers["Authorization"] = "Bearer " + os.Getenv("channelAccessToken")

	body, err := agent.DoReq()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("\n Body Reply : %+v\n", string(*body))

	return nil
}

func pushMessage(dat Data) error {

	mess1 := Message{
		Type: "text",
		Text: "Hai User ... ",
	}

	mess2 := Message{
		Type: "text",
		Text: "May I help you ...? ",
	}

	messages := []Message{mess1, mess2}

	rep := PushMessage{
		To:       "U772346mikhael73",
		Messages: messages,
	}

	agent := utils.NewHTTPRequest()
	agent.Url = "https://api.line.me"
	agent.Path = "/v2/bot/message/push"
	agent.Method = "POST"
	agent.IsJson = true
	agent.Json = rep

	agent.Headers["Authorization"] = "Bearer " + os.Getenv("channelAccessToken")

	body, err := agent.DoReq()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("\n Body Push Message : %+v\n", string(*body))

	return nil
}
