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

type Wisdom struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	apiUrl := "https://type.fit/api/quotes"

	quoteClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "krisraven-tms")

	res, err := quoteClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var w []Wisdom
 
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	randomnumber := rand.Intn(1643)

	// create a list/hashtable of all data. Use the random number as the key.
	// Then if that random number is in the list, output the quote
	counter := 1

	var quoteslist map[int]string
	quoteslist = make(map[int] string)

	for _, quotes := range w {
		// fmt.Printf("%d %s - %s\n", counter, quotes.Text, quotes.Author)
		quoteslist[counter]=(quotes.Text + " -" + quotes.Author)
		counter++
	}

	n := quoteslist[randomnumber]

	fmt.Printf("%+v \n", n)
	
}
