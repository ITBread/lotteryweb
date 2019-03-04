package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Http_Get(url string) string {
	//发送请求
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	err = json.Unmarshal(body, &ater)
}
