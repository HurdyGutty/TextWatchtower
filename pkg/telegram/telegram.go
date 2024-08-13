package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var Token string
var Chat_ID string

func getURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", Token)
}

func SendMessage(time time.Time, message string) (bool, error) {
	var err error
	var response *http.Response
	url := getURL()

	text := time.Format("17:00:00 24/02/2016") + ":" + message

	body, _ := json.Marshal(map[string]string{
		"chat_id": Chat_ID,
		"text":    text,
	})

	response, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}

	defer response.Body.Close()
	return true, nil
}
