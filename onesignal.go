package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://onesignal.com/api/v1/notifications"

	post := "{\"app_id\": \"9c8d3f26-33c1-4874-b179-7e416c006000\"," +
		"\"included_segments\": [\"All\"]," +
		"\"data\": {\"foo\": \"bar\"}," +
		"\"headings\": {\"en\": \"你好\"}," +
		"\"contents\": {\"en\": \"你\"}}"

	var jsonStr = []byte(post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Your Key")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
