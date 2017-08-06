package poloniex

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewPoloniexPublicAPI(client *http.Client, baseURL string) *PoloniexPublicAPI {
	return &PoloniexPublicAPI{
		client:  client,
		baseURL: baseURL,
	}
}

type PoloniexPublicAPI struct {
	client  *http.Client
	baseURL string
}

func (p *PoloniexPublicAPI) do(ctx context.Context, command string) ([]byte, error) {
	url := fmt.Sprintf("%s/public?command=%s", p.baseURL, command)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	log.Print(string(bs))
	return bs, err
}

func (p *PoloniexPublicAPI) ReturnCurrencies(ctx context.Context) (map[Currency]*CurrencyInfo, error) {
	bs, err := p.do(ctx, "returnCurrencies")
	if err != nil {
		return nil, err
	}

	var m map[Currency]*CurrencyInfo
	if err := json.Unmarshal(bs, &m); err != nil {
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

func (p *PoloniexPublicAPI) ReturnTicker(ctx context.Context) (map[CurrencyPair]*Ticker, error) {
	bs, err := p.do(ctx, "returnTicker")
	if err != nil {
		return nil, err
	}

	var m map[string]*Ticker
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, err
	}

	mm := make(map[CurrencyPair]*Ticker, len(m))
	for k, v := range m {
		cp, err := parseCurrencyPair(k)
		if err != nil {
			return nil, err
		}
		mm[cp] = v
	}

	return mm, nil
}
