// chuck is an Internet Chuck Norris DB client in Go.
// Usage: chuck [-n]
//   -n Only nerdy Chuck Norris jokes.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

// Response is the API response to a GET at http://api.icndb.com/jokes/random
type Response struct {
	Type  string
	Value Joke
}

// Joke is the actual Chuck Norris joke.
type Joke struct {
	Id         int
	Joke       string
	Categories []string
}

func main() {
	icndbURL := `http://api.icndb.com/jokes/random`
	nerdy := flag.Bool("n", false, "Only nerdy Chuck Norris jokes.")
	flag.Parse()
	if *nerdy {
		icndbURL += `?limitTo=[nerdy]`
	}
	resp, err := http.Get(icndbURL)
	if err != nil {
		log.Fatalf("Error connecting to API: %v", err)
	}
	defer resp.Body.Close()

	cnResp := new(Response)
	err = json.NewDecoder(resp.Body).Decode(cnResp)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	if cnResp.Type != "success" {
		log.Fatal("Error with API server: %s", cnResp.Type)
	}

	fmt.Println(html.UnescapeString(cnResp.Value.Joke))
}
