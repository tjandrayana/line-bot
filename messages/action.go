package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/anaskhan96/soup"
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

	// if err := SendMessages(dat, messages); err != nil {
	// 	log.Println("Push Message ERROR = ", err)
	// }

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

	receiver := []string{wawan}
	rep := PushMulticastMessage{
		To:       receiver,
		Messages: messages,
	}

	agent := utils.NewHTTPRequest()
	agent.Url = "https://api.line.me"
	// agent.Path = "/v2/bot/message/push"
	agent.Path = "/v2/bot/message/multicast"
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
	grady string = "U8b8591bd75661e51f1a80f74f92d2734"
)

func Horoskop() string {

	horoskop := ""

	resp, err := soup.Get("https://www.astrology.com/horoscope/daily/today/sagittarius.html")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "daily-horoscope").FindAll("p")

	for _, link := range links {
		fmt.Println(link.Text())
		horoskop = fmt.Sprintf("\n%s\n%s", horoskop, link.Text())
	}

	return horoskop
	// fmt.Printf("\n%+v\n", links.Data)
}
