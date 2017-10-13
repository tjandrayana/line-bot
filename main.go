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

// func triger(c *gin.Context) {

// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		log.Println("Error : ", err)
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": string(body),
// 	})

// }

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
