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

		fmt.Printf("\nMessages : %s\n", msg)

		if flag {
			result1, _ = gt.Translate(msg, "in", "en")
			result1 = strings.ToLower(result1)
			similarity := checkSimilarity2(msg, result1)
			fmt.Println("\nResult1 = ", result1, "\nSimilarity : ", similarity)

			if similarity < 0.85 {
				if result1 != msg {
					reply = fmt.Sprintf("%s, In English '%s' \nMeans : \n'%s'", namaUser, msg, result1)
					flag = false
				}
			}

		}

		if flag {
			result2, _ = gt.Translate(msg, "en", "in")
			result2 = strings.ToLower(result2)

			similarity := checkSimilarity2(msg, result2)
			fmt.Println("\nResult2 = ", result2, "\nSimilarity : ", similarity)
			if similarity < 0.85 {
				if result2 != msg {
					reply = fmt.Sprintf("%s, In Bahasa '%s' \nMeans : \n'%s'", namaUser, msg, result2)
					flag = false
				}
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

func checkSimilarity(messages, messages2 string) float64 {
	var similarity float64

	arrMessages1 := strings.Split(messages, " ")
	arrMessages2 := strings.Split(messages2, " ")

	var count int

	for i := 0; i < len(arrMessages2); i++ {
		for j := i; j < len(arrMessages1); j++ {
			if arrMessages2[i] == arrMessages1[j] {
				count++
			}
		}
	}

	fmt.Println("count = ", count)

	similarity = float64(count / len(arrMessages1))

	return similarity
}

func checkSimilarity2(messages, messages2 string) float64 {
	var similarity float64

	arrMessages1 := strings.Split(messages, " ")
	arrMessages2 := strings.Split(messages2, " ")

	map1 := make(map[string]bool)
	map2 := make(map[string]bool)
	globalMap := make(map[string]bool)

	for i := range arrMessages1 {
		if arrMessages1[i] == " " || arrMessages1[i] == "" {
			continue
		}
		map1[arrMessages1[i]] = true

		globalMap[arrMessages1[i]] = true
	}
	for i := range arrMessages2 {
		if arrMessages2[i] == " " || arrMessages2[i] == "" {
			continue
		}
		map2[arrMessages2[i]] = true

		globalMap[arrMessages2[i]] = true
	}

	var count int

	for key, _ := range map1 {
		for key2, _ := range map2 {
			if key == key2 {
				count++
			}
		}
	}

	fmt.Println("count = ", count, "\tlenmap2 = ", len(map2))
	similarity = float64(float64(count) / float64(len(map2)))

	return similarity
}
