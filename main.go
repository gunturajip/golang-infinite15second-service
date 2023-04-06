package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	randomWaterWind()

	ticker := time.NewTicker(15 * time.Second)

	// Creating channel using make
	tickerChan := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChan:
				return
			// interval task
			case <-ticker.C:
				randomWaterWind()
			}
		}
	}()

	for {
	}
	// // Calling Sleep() method
	// time.Sleep(100 * time.Second)

	// // Calling Stop() method
	// ticker.Stop()

	// // Setting the value of channel
	// tickerChan <- true

	// // Printed when the ticker is turned off
	// fmt.Println("Ticker is turned off!")
}

func randomWaterWind() {
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// print hasil convert response slice of byte ke json
	type Response struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	}
	var response Response
	_ = json.Unmarshal([]byte(body), &response)
	exp, _ := json.MarshalIndent(Response{Water: response.Water, Wind: response.Wind}, "", "  ")
	fmt.Println(string(exp))

	// print hasil status dari response
	var status_water string
	var status_wind string
	if response.Water < 5 {
		status_water = "aman"
	} else if response.Water >= 6 && response.Water <= 8 {
		status_water = "siaga"
	} else if response.Water > 8 {
		status_water = "bahaya"
	}
	if response.Wind < 6 {
		status_wind = "aman"
	} else if response.Wind >= 7 && response.Wind <= 15 {
		status_wind = "siaga"
	} else if response.Wind > 15 {
		status_wind = "bahaya"
	}
	fmt.Printf("status water : %s\n", status_water)
	fmt.Printf("status wind : %s\n", status_wind)
}
