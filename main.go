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
	"strings"

	gt "github.com/bas24/googletranslatefree"

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

	// const text string = `buku`
	// result, _ := gt.Translate(text, "in", "eng")
	// fmt.Println(result)

	User = make(map[string]string)

	User[fery] = "Fery"
	User[wawan] = "Wawan"
	User[dwicky] = "Dwicky"
	User[yosia] = "Yosia"

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

	messages := checkMessage(dat)

	if err := reply(dat, messages); err != nil {
		log.Println("Reply ERROR = ", err)
	}

	if err := pushMessage(dat, messages); err != nil {
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

func reply(dat Data, messages []Message) error {

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

func pushMessage(dat Data, messages []Message) error {

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
	wawan  string = "U0d7ba35d0e9e44f209d37f9bdf81d2b9"
	dwicky string = "U4d3ecc4048a8e14040f28af321c089ef"
	fery   string = "Ud92705b0a3468dc19885afea426ad219"
	yosia  string = "Ud71e2b65ef7569017bc9f32712cee4d5"
)

var User map[string]string

func checkMessage(dat Data) []Message {

	var messages []Message

	if dat.Events[0].Type == "follow" {
		mess1 := Message{
			Type: "text",
			Text: "Thx ya sudah add aku sebagai teman kamu.",
		}

		mess2 := Message{
			Type: "text",
			Text: "Perkenalkan nama saya Hero, saya adalah bot chat yang sedang dikembangkan ...",
		}
		messages = append(messages, mess1)
		messages = append(messages, mess2)

	} else {
		var reply string

		msg := strings.ToLower(dat.Events[0].Message.Text)

		result, _ := gt.Translate(msg, "in", "eng")
		fmt.Println(result)

		namaUser := User[dat.Events[0].Source.UserID]
		if namaUser == "" {
			namaUser = "hei "
		}

		reply = fmt.Sprintf("%s, In english '%s' \nmeans : \n'%s'", namaUser, msg, result)

		mess1 := Message{
			Type: "text",
			Text: reply,
		}

		messages = append(messages, mess1)
	}

	return messages

}

var CollectionBahasaGaul map[string]string

func fillCollection() {
	CollectionBahasaGaul = make(map[string]string)
}

type Means struct {
	Word        string `json:"word"`
	Definitions []struct {
		Definition   string `json:"definition"`
		PartOfSpeech string `json:"partOfSpeech"`
	} `json:"definitions"`
}

func DefinitionWord(word string) (Means, error) {

	var means Means

	agent := utils.NewHTTPRequest()
	agent.Url = "http://dictionary.cambridge.org/"
	agent.Path = "dictionary/indonesian-english" + word
	agent.Method = "GET"

	body, err := agent.DoReq()
	if err != nil {
		log.Println(err)
		return means, err
	}

	// if err := json.Unmarshal(*body, &means); err != nil {
	// 	log.Println(err)
	// 	return means, err
	// }

	fmt.Printf("\n BODY = %+v\n", string(*body))

	return means, nil
}
