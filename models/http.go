package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (body []byte, err error) {
	fmt.Println(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	// reqest.Header.Set("Cookie", "appver=1.5.0.75771")
	// reqest.Header.Set("Referer", "http://music.163.com/")

	response, err := client.Do(reqest)

	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
	}
	return
}

func HttpPost(url, data string) (body []byte, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return
	}

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")

	response, err := client.Do(reqest)

	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
	}
	return
}
