package messages

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	gt "github.com/bas24/googletranslatefree"
)

func CheckMessage(dat Data) []Message {
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
		var reply, result1, result2 string
		flag := true
		namaUser := "hei "

		msg := strings.ToLower(dat.Events[0].Message.Text)

		if last := len(msg) - 1; last >= 0 && msg[last] == ' ' {
			msg = msg[:last]
		}

		if flag {
			result1, _ = gt.Translate(msg, "in", "en")
			result1 = strings.ToLower(result1)
			fmt.Println(result1)
			if result1 != msg {
				reply = fmt.Sprintf("%s, In english '%s' \nmeans : \n'%s'", namaUser, msg, result1)
				flag = false
			}

		}

		if flag {
			result2, _ = gt.Translate(msg, "en", "in")
			result2 = strings.ToLower(result2)
			fmt.Println(result2)
			if result2 != msg {
				reply = fmt.Sprintf("%s, In Indonesian '%s' \nmeans : \n'%s'", namaUser, msg, result2)
				flag = false
			}
		}

		if flag {
			result1, _ = gt.Translate(msg, "in", "en")
			result1 = strings.ToLower(result1)

			reply = namaUser + ", "
			reply = fmt.Sprintf("%s In english '%s' \nmeans : \n'%s'\n", reply, msg, result1)

			result2, _ = gt.Translate(msg, "en", "in")
			result2 = strings.ToLower(result2)

			reply = fmt.Sprintf("\n%s AND \nIn Indonesian '%s' \nmeans : \n'%s'", reply, msg, result2)

		}

		mess1 := Message{
			Type: "text",
			Text: reply,
		}

		messages = append(messages, mess1)
	}

	return messages
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
