package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendHttpRequest() {

	url := "https://google.com.br"

	response, requestError := http.Get(url)

	if requestError != nil {
		fmt.Println(requestError)
	}

	value, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		fmt.Println(readError)
	}

	fmt.Println(string(value))

}

func main() {

	fmt.Println("Send HTTP request")

	sendHttpRequest()

}
