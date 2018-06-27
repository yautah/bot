package network

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	. "github.com/yautah/bot/data"
	"log"
	"time"
)

const (
	auth     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzAyLCJpZGVuIjoiNDE5MDU2NjY2NTcxNTcxMjIwIiwibWQiOnt9LCJ0cyI6MTUyNzc1MjkyNjM0M30.Uf3RPtgOkmSEUEXeaqz3HdudbMPhLGlp9hlH1Rq4nhQ"
	endpoint = "http://api.royaleapi.com"
)

type filtertype func(r Result) bool

func sliceFilter(slice []Result, a filtertype) (ftslice []Result) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

func FetchWars(tag string) []Result {
	resp, err := resty.
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(20)).
		R().
		SetQueryParams(map[string]string{
			"max": "1",
			// "page": "1",
		}).
		SetHeader("auth", auth).
		Get(endpoint + fmt.Sprintf("/clan/%s/battles?type=war", tag))

	if err != nil {
		log.Fatal(err)
	}

	results := make([]Result, 0)

	err = json.Unmarshal(resp.Body(), &results)

	if err != nil {
		log.Fatal(err)
	}

	var warFilter filtertype = func(r Result) bool {
		// return r.WarType != "clanWarWarDay"
		return true
	}

	var timeFilter filtertype = func(r Result) bool {
		// cur := time.Now()
		// timestamp := cur.UnixNano() / 1000000000
		// log.Printf("---------filter time: %s vs %s, %d", r.Team[0].Name, r.Opponent[0].Name, timestamp-r.UtcTime)
		// return timestamp-r.UtcTime < 5*60
		return true
	}

	log.Printf("---------results: %d", len(results))
	results = sliceFilter(sliceFilter(results, warFilter), timeFilter)
	log.Printf("---------results: %d", len(results))
	log.Printf("---------results: %d", time.Now().UnixNano()/1000000000)
	for _, v := range results {
		log.Printf("---------result: %s, %d", v.WarType, v.UtcTime)
	}

	return results
}

func FetchChest(tag string) Chest {
	resp, err := resty.
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(20)).
		R().
		SetHeader("auth", auth).
		Get(endpoint + fmt.Sprintf("/player/%s/chests", tag))

	if err != nil {
		log.Fatal(err)
	}

	chest := Chest{}
	err = json.Unmarshal(resp.Body(), &chest)

	if err != nil {
		log.Fatal(err)
	}
	return chest
}
