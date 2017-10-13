package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
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
	var channelSecret string
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("error : ", err)
		return
	}

	fmt.Printf("\n%+v\n", string(body))

	c.JSON(200, gin.H{
		"message": string(body),
	})

	if !validateSignature(channelSecret, c.Request.Header.Get("X-Line-Signature"), body) {
		log.Println("error : ", err)
		return
	}

	c.JSON(200, gin.H{
		"message": string(body),
	})

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
