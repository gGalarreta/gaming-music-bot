package facebook

import (
	"github.com/tidwall/gjson"
	"net/http"
	"bytes"
	"os"
)

func GetMessageData(message string) ( sender_id string, text string, postback string)  {

	data_sender := gjson.Get(message, "entry.#.messaging.#.sender.id")
	data_text := gjson.Get(message, "entry.#.messaging.#.message")
	data_postback := gjson.Get(message, "entry.#.messaging.#.postback")
	sender_id = data_sender.Array()[0].Array()[0].String()
	text = HandleMessage(data_text.Array()[0].Array()[0].String())
	postback = HandlePostback(data_postback.Array()[0].String())
	return
}

func HandleMessage(message string) (text string) {
	text = ""
	data_text := gjson.Get(message, "text").String()
	if data_text != "" {
		text = data_text
	}
	return
}

func HandlePostback(message string) (postback string)  {
	postback = ""
	if len(message) > 0 {
		postback = gjson.Get(message, "postback").String()
	}
	return
}

func SendingText(sender string, message string) {
    jsonStr := []byte(`{	"recipient":{
		"id":` + sender + `
	  },
	  "message":{
		"text": ` + message + `
	  }}`)
	  PostData(jsonStr)
}

func PostData(post_content []byte)  {
    req, err := http.NewRequest("POST", FacebbokUrl(), bytes.NewBuffer(post_content))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

func FacebbokUrl() string {
	return os.Getenv("FB_URL") + os.Getenv("FB_TOKEN")
}