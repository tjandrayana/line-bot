package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tjandrayana/line-bot/utils"
)

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

	messages := CheckMessage(dat)

	if err := ReplyMessages(dat, messages); err != nil {
		log.Println("Reply ERROR = ", err)
	}

	if err := SendMessages(dat, messages); err != nil {
		log.Println("Push Message ERROR = ", err)
	}

	fmt.Println("\nSuccess\n")

}

func ReplyMessages(dat Data, messages []Message) error {

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

	_, err := agent.DoReq()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SendMessages(dat Data, messages []Message) error {

	rep := PushMessage{
		To:       wawan,
		Messages: messages,
	}

	agent := utils.NewHTTPRequest()
	agent.Url = "https://api.line.me"
	agent.Path = "/v2/bot/message/push"
	agent.Method = "POST"
	agent.IsJson = true
	agent.Json = rep

	agent.Headers["Authorization"] = "Bearer " + os.Getenv("channelAccessToken")

	_, err := agent.DoReq()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

const (
	wawan string = "U0d7ba35d0e9e44f209d37f9bdf81d2b9"
)
