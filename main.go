package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"math/rand"
	"time"
	"fmt"
)

type Wisdom[] struct {
	Text  string `json:"text"`
	Author string `json:"author"`
}

type QuoteMap struct {
	n string
}

func main() {
	apiUrl := "https://type.fit/api/quotes"

	qq, err := getQuotes(apiUrl)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	randomnumber := rand.Intn(1643)

	counter := 1

	var quotesList map[int]string
	quotesList = make(map[int]string)

	for _, quotes := range qq {
		quotesList[counter]=(quotes.Text + " -" + quotes.Author)
		counter++
	}

	n := quotesList[randomnumber]

	fmt.Printf("%+v \n", n)
}

func getQuotes(apiUrl string)(Wisdom, error) {
	w := Wisdom{}
	
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return w, err
	}

	req.Header.Set("User-Agent", "krisraven-tms")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return w, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return w, err
	}

	err = json.Unmarshal(body, &w)
	if err != nil {
		return w, err
	}

	return w, err
}