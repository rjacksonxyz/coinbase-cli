package internal

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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

type CommandInfo struct {
	InfoType string
}

type CoinbaseClient struct {
	Creds   CoinbaseAPICredentials
	BaseUrl string
	Client  *http.Client
}

var baseUrl string = "https://api.coinbase.com/v2/"

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
		return CoinbaseClient{}, errors.New("unable to parse credentials from provied file/path")
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

func (c *CoinbaseClient) Get(uri string) ([]byte, error) {
	return c.SendRequest("GET", uri)
}

func (c *CoinbaseClient) SendRequest(method string, uri string) ([]byte, error) {

	req, err := http.NewRequest(method, c.BaseUrl+uri, nil)
	fmt.Println(c.BaseUrl + uri)
	if err != nil {
		log.Println("unable to create new request")
	}
	c.authenticate(req, method, uri)
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

// API Key + Secret authentication requires a request header of the HMAC SHA-256
// signature of the "message" as well as an incrementing nonce and the API key
func (c *CoinbaseClient) authenticate(req *http.Request, method string, uri string) error {

	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + method + uri + "" //As per Coinbase Documentation

	h := hmac.New(sha256.New, []byte(c.Creds.APISecret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("CB-ACCESS-KEY", c.Creds.APIKey)
	req.Header.Set("CB-ACCESS-SIGN", signature)
	req.Header.Set("CB-ACCESS-TIMESTAMP", nonce)
	req.Header.Set("CB-VERSION", "2021-01-01")

	return nil
}

func CoinbaseCLIRequest() error {

	return nil
}
