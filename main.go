package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	fmt.Println(h8HelperRand.RandomInt(10, 20))

	// get url request
	// siapkan get request dengan url yang diinginkan
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// convert response ke format tipe data slice of byte
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// print hasil convert response slice of byte ke string
	fmt.Println(string(responseBody))

	// post url request
	// siapkan data yang ingin dikirim
	data := map[string]interface{}{
		"water": h8HelperRand.RandomInt(1, 100),
		"wind":  h8HelperRand.RandomInt(1, 100),
	}

	// convert data ke tipe json
	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	// siapkan post request dengan url yang diinginkan
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	// kirim post request
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	// convert response ke format tipe data slice of byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// print hasil convert response slice of byte ke string
	fmt.Println(string(body))
}
