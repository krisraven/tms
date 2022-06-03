package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type Wisdom struct {
	Text   string
	Author string
}

func main() {
	apiUrl := "https://type.fit/api/quotes"

	quotes, err := getQuote(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v\n", quotes.Text)

	for _, w := range quotes {
		fmt.Printf("Quote by: %+v\n", w.Text)
	}
}

// func main() {
// 	apiUrl := "https://type.fit/api/quotes"

// 	quotes, err := getQuote(apiUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Printf("%d quotes\n", quotes.Count)

// }

func getQuote(apiUrl string) (Wisdom, error) {
	w := Wisdom{}
  	// var w []Wisdom

	// we have to check for errors, so we check for nil/null

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return w, err
	}

	req.Header.Set("User-Agent", "krisraven-tms-test")

	res, err := http.DefaultClient.Do(req) // make the request
	if err != nil {
		return w, err
	}

	if res.Body != nil {
		defer res.Body.Close() // defer is used to defer any cleanup activities until the end of the function
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return w, err
	}

	// Unmarshal takes the response body and tries to input the data in the data structure created in a struct
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatalf("unable to parse value: %q, error: %s",
			string(body), err.Error()) // error handling
		return w, err
	}

	return w, nil
}

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "log"
//     "net/http"
// )

// type Wisdom struct {
// 	Text   string
// 	Author string
// }

// func main() {
//     apiUrl := "https://type.fit/api/quotes"
//     res, err := http.Get(apiUrl)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer res.Body.Close()
//     body, err := ioutil.ReadAll(res.Body)
//     if err != nil {
//         log.Fatal(err)
//     }

//     var w []Wisdom

//     err = json.Unmarshal(body, &w)
//     if err != nil {
//         panic(err)
//     }

//     for _, quotes := range w {
//         log.Fatal(fmt.Println("\"",quotes.Text,"\""," -", quotes.Author))
//     }
// }