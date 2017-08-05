package poloniex

import (
	"github.com/pkg/errors"
	"strings"
)

type CurrencyPair struct {
	Trading    Currency
	Settlement Currency
}

func NewCurrencyPair(trading Currency, settlement Currency) CurrencyPair {
	return CurrencyPair{
		Trading:    trading,
		Settlement: settlement,
	}
}

func parseCurrencyPair(str string) (CurrencyPair, error) {
	xs := strings.Split(str, "_")
	if len(xs) != 2 {
		return CurrencyPair{}, errors.Errorf("cannot parse currency pair '%s'", str)
	}
	return NewCurrencyPair(Currency(xs[1]), Currency(xs[0])), nil
}

type Currency string

type CurrencyInfo struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	MaxDailyWithdrawal int     `json:"maxDailyWithdrawal"`
	TxFee              float64 `json:"txFee,string"`
	MinConf            int     `json:"minConf"`
	Disabled           int     `json:"disabled"`
	DepositAddress     string  `json:"depositAddress"`
	Delisted           int     `json:"delisted"`
	Frozen             int     `json:"frozen"`
}
