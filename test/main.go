package main

import (
	"context"
	"encoding/json"
	"github.com/tattsun/poloniex"
	"log"
	"net/http"
)

func main() {
	polo := poloniex.NewPoloniexPublicAPI(http.DefaultClient, "https://poloniex.com")
	ticker, err := polo.ReturnTicker(context.Background())
	if err != nil {
		panic(err)
	}
	for k, _ := range ticker {
		bs, err := json.Marshal(k)
		if err != nil {
			panic(err)
		}
		log.Print(string(bs))
	}
	bs, err := json.MarshalIndent(ticker, "", "\t")
	//	bs, err := json.Marshal(ticker)
	if err != nil {
		panic(err)
	}
	log.Print(string(bs))
}
