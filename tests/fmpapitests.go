package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://financialmodelingprep.com/api/v3/profile/NVDA?apikey=439590488dcfe01c49ae5bb9f3213be1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString :=string(responseData)

	fmt.Println(responseString)
}


