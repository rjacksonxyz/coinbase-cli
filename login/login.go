package login

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fabioberger/coinbase-go"
)

type CoinbaseAPICredentials struct {
	APIKey    string `json:"API-Key"`
	APISecret string `json:"API-Secret"`
}

func ClientFromJSON(filepath string) coinbase.Client {
	raw, _ := ioutil.ReadFile(filepath)
	var creds CoinbaseAPICredentials
	err := json.Unmarshal(raw, &creds)
	if err != nil {
		log.Println("unable to parse credentials from provied file/path")
	}
	log.Println(creds)
	return coinbase.ApiKeyClient(creds.APIKey, creds.APISecret)
}
