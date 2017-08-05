package poloniex

type Currency struct {
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

type Currencies struct {
	m map[string]Currency
}

func newCurrencies(m map[string]Currency) *Currencies {
	return &Currencies{
		m: m,
	}
}

func (c *Currencies) Map() map[string]Currency {
	nm := make(map[string]Currency, len(c.m))
	for k, v := range c.m {
		nm[k] = v
	}
	return nm
}

func (c *Currencies) Get(key string) (Currency, bool) {
	if c, ok := c.m[key]; ok {
		return c, true
	}
	return Currency{}, false
}

func (c *Currencies) put(key string, cur Currency) {
	c.m[key] = cur
}
