package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func externalCall(client *http.Client) string {

	var bodyBytes []byte
	n := rand.Int() % len(apis)
	// fmt.Println("Gonna work from home...", apis[n])
	request, err := http.NewRequest("GET", apis[n], nil)
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, _ = ioutil.ReadAll(response.Body)
	}
	bodyString := string(bodyBytes)

	return bodyString
}
