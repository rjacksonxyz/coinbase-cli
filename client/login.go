package login

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CoinbaseAPICredentials struct {
	APIKey    string `json:"API-Key"`
	APISecret string `json:"API-Secret"`
}

type CoinbaseClient struct {
	Creds   CoinbaseAPICredentials
	BaseUrl string
	Client  *http.Client
}

func NewAPIClient(creds CoinbaseAPICredentials) CoinbaseClient {
	c := CoinbaseClient{
		Creds:   creds,
		BaseUrl: "https://api.coinbase.com",
		Client:  &http.Client{},
	}
	return c
}

func ClientFromJSON(filepath string) CoinbaseClient {
	raw, _ := ioutil.ReadFile(filepath)
	var creds CoinbaseAPICredentials
	err := json.Unmarshal(raw, &creds)
	if err != nil {
		log.Println("unable to parse credentials from provied file/path")
	}
	c := NewAPIClient(creds)
	return c
}

func (c CoinbaseClient) Get() map[string]interface{} {

	req, err := http.NewRequest("GET", c.BaseUrl, nil)
	if err != nil {
		log.Println("unable to create new request")
	}
	c.authenticate(req)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Print(err)
	}
	log.Println("Status: ", resp.Status)
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bytes := buf.Bytes()
	log.Print(string(bytes))

	var data map[string]interface{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Println("unable to parse response")
	}

	return data
}

// API Key + Secret authentication requires a request header of the HMAC SHA-256
// signature of the "message" as well as an incrementing nonce and the API key
func (c CoinbaseClient) authenticate(req *http.Request) error {

	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + "GET" + "/v2/accounts" + "" //As per Coinbase Documentation

	h := hmac.New(sha256.New, []byte(c.Creds.APISecret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("CB-ACCESS-KEY", c.Creds.APIKey)
	req.Header.Set("CB-ACCESS-SIGN", signature)
	req.Header.Set("CB-ACCESS-TIMESTAMP", nonce)

	return nil
}
