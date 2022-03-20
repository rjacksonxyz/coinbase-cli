package internal

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

type NewRequestInfo struct {
	Method string
	Uri    string
}

type CommandInfo struct {
	InfoType string
}

type CoinbaseClient struct {
	Creds   CoinbaseAPICredentials
	BaseUrl string
	Client  *http.Client
}

var baseUrl string = "https://api.coinbase.com/v2/"
var priceUri string = "prices/"

// ClientFromJSON returns a CoinbaseClient given a filepath to a json file.
/* Example JSON:
{
	"API-Key": "<insert API Key>",
	"API-Secret" : "<insert API Secret>"
}
*/
func NewClient(filepath string) (CoinbaseClient, error) {
	raw, _ := ioutil.ReadFile(filepath)
	var creds CoinbaseAPICredentials
	err := json.Unmarshal(raw, &creds)
	if err != nil {
		fmt.Println("unable to parse credentials from provied file/path")
		return CoinbaseClient{}, err
	}
	c := ClientFromCredentials(creds)
	return c, nil
}

// NewAPI client returns a CoinbaseClient that holds credentials and a base url
// used to execute requests to the Coinbase API.
func ClientFromCredentials(creds CoinbaseAPICredentials) CoinbaseClient {
	c := CoinbaseClient{
		Creds:   creds,
		BaseUrl: baseUrl,
		Client:  &http.Client{},
	}
	return c
}

func (c *CoinbaseClient) Get(reqData NewRequestInfo) ([]byte, error) {
	return c.SendRequest("GET", reqData)
}

func (c *CoinbaseClient) SendRequest(method string, reqData NewRequestInfo) ([]byte, error) {

	req, err := http.NewRequest(method, c.BaseUrl, nil)
	if err != nil {
		log.Println("unable to create new request")
	}
	c.authenticate(req, reqData)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bytes := buf.Bytes()

	var data map[string]interface{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("unable to parse response body")
	}

	o, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}

	return o, nil
}

// func (c *CoinbaseClient) GetAccounts() ([]byte, error) {

// 	info := NewRequestInfo{Method: "GET", Uri: "/v2/accounts?&limit=100&order=asc"}
// 	data, err := c.Get(info)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// API Key + Secret authentication requires a request header of the HMAC SHA-256
// signature of the "message" as well as an incrementing nonce and the API key
func (c *CoinbaseClient) authenticate(req *http.Request, reqData NewRequestInfo) error {

	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + reqData.Method + reqData.Uri + "" //As per Coinbase Documentation

	h := hmac.New(sha256.New, []byte(c.Creds.APISecret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("CB-ACCESS-KEY", c.Creds.APIKey)
	req.Header.Set("CB-ACCESS-SIGN", signature)
	req.Header.Set("CB-ACCESS-TIMESTAMP", nonce)
	req.Header.Set("CB-VERSION", "2021-01-01")

	return nil
}

func parseCommand() {

}

func CoinbaseCLIRequest() error {
	return nil
}
