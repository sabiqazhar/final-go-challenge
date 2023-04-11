package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	ticktack := time.NewTicker(15 * time.Second)
	for s := range ticktack.C {
		fmt.Println(s)
		fmt.Println(strings.Repeat("#", 25))
		placeHolder()
		randNum()
	}

}

func placeHolder() {
	randUser := rand.Intn(10)
	data := map[string]interface{}{
		"title": "darwin",
		"body" : "hawkins",
		"userId": randUser,
	}

	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

func randNum()  {
	var (
		randWater = rand.Intn(100)
		randWind = rand.Intn(100)
		statusWater = "aman"
		statusWind = "aman"
	)

	type field struct {
		Water int `json:"water"`
		Wind int `json:"wind"`
	}

	if randWater >= 6 && randWater <= 8 {
		statusWater = "siaga"
	} else if randWater > 8 {
		statusWater = "bahaya"
	}

	if randWind >= 7 && randWind <= 15 {
		statusWind = "siaga"
	} else if randWater > 8 {
		statusWind = "bahaya"
	}

	jsonString := field{
		Water: randWater,
		Wind: randWind,
	}

	res, err := json.MarshalIndent(jsonString, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
	fmt.Println()
	fmt.Println("status water: ", statusWater)
	fmt.Println("status wind: ", statusWind)
}