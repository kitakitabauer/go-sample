package main

import (
	"net/url"
	"net/http"
	"strings"
	"io/ioutil"
)

func main() {
	client := &http.Client{}

	put(client)
	delete(client)
}

func put(client *http.Client) {
	data := url.Values{"foo": {"bar"}}

	req, _ := http.NewRequest(
		"PUT",
		"http://httpbin.org/put",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlenccoded")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}

func delete(client *http.Client) {
	data := url.Values{"foo": {"bar"}}

	req, _ := http.NewRequest(
		"DELETE",
		"http://httpbin.org/delete",
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlenccoded")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
