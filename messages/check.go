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
			Text: fmt.Sprintf("Terima kasih sudah menambahkan aku sebagai teman kamu.\nPerkenalkan namaku Hero, Aku adalah bot chat yang sedang dikembangkan ...\nAku akan mencoba membantumu dalam berkomunikasi\nSilahkan kamu tulis kata dalam Bahasa Indonesia atau Inggris kemudian aku akan mengartikan untuk anda\n"),
		}

		mess2 := Message{
			Type: "text",
			Text: fmt.Sprintf("Thank you for adding me as your friend.\nIntroduce my name Hero, I am a chat bot being developed.\nI will try to help you in communicating.\nPlease write the word in Indonesian or English then I will interpret for you.\n"),
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
				reply = fmt.Sprintf("%s, In English '%s' \nMeans : \n'%s'", namaUser, msg, result1)
				flag = false
			}

		}

		if flag {
			result2, _ = gt.Translate(msg, "en", "in")
			result2 = strings.ToLower(result2)
			fmt.Println(result2)
			if result2 != msg {
				reply = fmt.Sprintf("%s, In Bahasa '%s' \nMeans : \n'%s'", namaUser, msg, result2)
				flag = false
			}
		}

		if flag {
			result1, _ = gt.Translate(msg, "in", "en")
			result1 = strings.ToLower(result1)

			reply = namaUser + ", "
			reply = fmt.Sprintf("%s In English '%s' \nMeans : \n'%s'\n", reply, msg, result1)

			result2, _ = gt.Translate(msg, "en", "in")
			result2 = strings.ToLower(result2)

			reply = fmt.Sprintf("\n%s AND \nIn Bahasa '%s' \nMeans : \n'%s'", reply, msg, result2)

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
