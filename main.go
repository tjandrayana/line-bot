package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
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

	var botRequest BotRequest
	if c.Bind(&botRequest) == nil {
		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("request convert error, request data is %+v", botRequest)})
		return
	}

	for _, result := range botRequest.Result {

		request := SendRequest{
			To:        []string{result.Content.From},
			ToChannel: ToChannel,
			EventType: EventType,
			Content: Content{
				ContentType: result.Content.ContentType,
				ToType:      result.Content.ToType,
				Text:        result.Content.Text,
			},
		}

		if _, err := post(request); err != nil {
			log.Printf("Error: %s", err.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "end"})

}

type BotRequestBody struct {
	Request BotRequest `json:"Request"`
}

//BotRequest
type BotRequest struct {
	Result []BotResult `json:"result"`
}

type BotResult struct {
	From        string   `json:"from"`
	FromChannel string   `json:"fromChannel"`
	To          []string `json:"to"`
	ToChannel   string   `json:"toChannel"`
	EventType   string   `json:"eventType"`
	ID          string   `json:"id"`
	Content     Content  `json:"content"`
}

type Content struct {
	ID          string   `json:"id"`
	ContentType int      `json:"contentType"`
	From        string   `json:"from"`
	CreatedTime int      `json:"createdTime"`
	To          []string `json:"to"`
	ToType      int      `json:"toType"`
	Text        string   `json:"text"`
}

type SendRequest struct {
	To        []string `json:"to"`
	ToChannel int      `json:"toChannel"`
	EventType string   `json:"eventType"`
	Content   Content  `json:"content"`
}

const (
	EndPoint  = "tjandrayana-line-bot.herokuapp.com/line"
	ToChannel = 1540625385
	EventType = "138311608800106203"
)

func callbackHandler(c *gin.Context) {

}

func post(r SendRequest) (*http.Response, error) {
	b, _ := json.Marshal(r)
	req, _ := http.NewRequest(
		"POST",
		EndPoint,
		bytes.NewBuffer(b),
	)

	req = setHeader(req)

	proxyURL, _ := url.Parse(os.Getenv("FIXIE_URL"))
	client := &http.Client{
		Timeout:   time.Duration(15 * time.Second),
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
	}

	return client.Do(req)
}

func setHeader(req *http.Request) *http.Request {
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Line-ChannelID", os.Getenv("LINE_CHANNEL_ID"))
	req.Header.Add("X-Line-ChannelSecret", os.Getenv("LINE_CHANNEL_SECRET"))
	req.Header.Add("X-Line-Trusted-User-With-ACL", os.Getenv("LINE_CHANNEL_MID"))
	return req
}
