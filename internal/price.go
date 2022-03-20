package internal

import (
	"fmt"
	"strings"
)

var priceUri string = "prices/"

type PriceRequest struct {
	CurrencyPairs string
}

func (c *CoinbaseClient) GetPrice(params PriceRequest) ([]byte, error) {
	for _, pair := range strings.Split(params.CurrencyPairs, ",") {
		if data, err := c.SendRequest("GET", constructPriceURI(pair)); err != nil {
			return nil, err
		} else {
			fmt.Println(string(data))
		}
	}
	return nil, nil
	//return c.SendRequest("GET", reqData)
}

func constructPriceURI(ticker string) string {
	return fmt.Sprintf("%s%s/spot", priceUri, ticker)
}
