package facebook

import (
	"net/http"
	"fmt"
	"bytes"
	"os"
)

func SendingText() {
    var jsonStr = []byte(`{	"recipient":{
		"id":"53844"
	  },
	  "message":{
		"text":"hello, world!"
	  }}`)
    req, err := http.NewRequest("POST", FacebbokUrl(), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
}

func FacebbokUrl() string {
	return "https://graph.facebook.com/v2.6/me/messages?access_token=" + os.Getenv("FB_TOKEN")
}