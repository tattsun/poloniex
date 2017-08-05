package poloniex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PoloniexPublicAPI struct {
	baseURL string
}

func (p *PoloniexPublicAPI) ReturnCurrencies(ctx context.Context) (map[Currency]*CurrencyInfo, error) {
	cli := http.DefaultClient

	url := fmt.Sprintf("%s/public?command=returnCurrencies", p.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	var m map[Currency]*CurrencyInfo
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}

type Tickers map[*CurrencyPair]*Ticker

type Ticker struct {
	Last          float64 `json:"last,string"`
	LowestAsk     float64 `json:"lowestAsk,string"`
	HighestBid    float64 `json:"highestBid,string"`
	PercentChange float64 `json:"percentChange,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
}

func (p *PoloniexPublicAPI) ReturnTicker(ctx context.Context) map[CurrencyPair]*Ticker {
	return nil
}
