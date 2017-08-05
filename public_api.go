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

func (p *PoloniexPublicAPI) ReturnCurrencies(ctx context.Context) (*Currencies, error) {
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

	var m map[string]Currency
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}

	cs := newCurrencies(m)
	return cs, nil
}
