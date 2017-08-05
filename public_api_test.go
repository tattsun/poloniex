package poloniex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func strAddr(str string) *string {
	return &str
}

func TestReturnCurrencies(t *testing.T) {
	m := map[string]Currency{
		"1CR": Currency{
			ID:             1,
			Name:           "1CRedit",
			TxFee:          0.01,
			MinConf:        3,
			DepositAddress: "",
			Disabled:       0,
			Delisted:       1,
			Frozen:         0,
		},
		"ABY": Currency{
			ID:             2,
			Name:           "ArtByte",
			TxFee:          0.12340001,
			MinConf:        8,
			DepositAddress: "abcdefg",
			Disabled:       0,
			Delisted:       1,
			Frozen:         0,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"1CR":{"id":1,"name":"1CRedit","txFee":"0.01000000","minConf":3,"depositAddress":null,"disabled":0,"delisted":1,"frozen":0},"ABY":{"id":2,"name":"ArtByte","txFee":"0.12340001","minConf":8,"depositAddress":"abcdefg","disabled":0,"delisted":1,"frozen":0}}`)
	}))
	defer ts.Close()

	cli := &PoloniexPublicAPI{ts.URL}
	cs, err := cli.ReturnCurrencies(context.Background())
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(cs.Map(), m) {
		t.Fatal(cmp.Diff(m, cs.Map(), nil))
	}
}
