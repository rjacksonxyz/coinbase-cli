package client

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type CoinbaseClient struct {
	Creds   CoinbaseAPICredentials
	BaseUrl string
	Client  *http.Client
}

//TODO: refactor to take in two params for credentials
// NewAPI client returns a CoinbaseClient that holds credentials and a base url
// used to execute requests to the Coinbase API.
func NewAPIClient(creds CoinbaseAPICredentials) CoinbaseClient {
	c := CoinbaseClient{
		Creds:   creds,
		BaseUrl: "https://api.coinbase.com/v2/accounts?&limit=100&order=asc",
		Client:  &http.Client{},
	}
	return c
}

// ClientFromJSON returns a CoinbaseClient given a JSON file path.
/* Example JSON:

{
	"API-Key": "<insert API Key>",
	"API-Secret" : "<insert API Secret>"
}

*/
func ClientFromJSON(filepath string) (CoinbaseClient, error) {
	raw, _ := ioutil.ReadFile(filepath)
	var creds CoinbaseAPICredentials
	err := json.Unmarshal(raw, &creds)
	if err != nil {
		fmt.Println("unable to parse credentials from provied file/path")
		return CoinbaseClient{}, err
	}
	c := NewAPIClient(creds)
	return c, nil
}

// ClientFromStdIn returns a CoinbaseClient given prompted inputs (api-key and api-secret)
func ClientFromStdIn() CoinbaseClient {

	var key string
	var secret string
	var err error

	login := false
	for !login {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("enter api-key: ")
		key, err = reader.ReadString('\n')
		if err == io.EOF {
			func() {}()
		} else if err != nil {
			fmt.Println("Encounterd error on input. Please try again")
			continue
		}

		fmt.Print("enter api-secret: ")
		secret, err = reader.ReadString('\n')
		if err == io.EOF {
			func() {}()
		} else if err != nil {
			fmt.Println("Encounterd error on input. Please try again")
			continue
		}
		login = true
	}
	creds := CoinbaseAPICredentials{APIKey: key, APISecret: secret}
	c := NewAPIClient(creds)
	return c
}

func (c *CoinbaseClient) Get(reqData NewRequestInfo) ([]byte, error) {

	req, err := http.NewRequest("GET", c.BaseUrl, nil)
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

func (c *CoinbaseClient) GetPrice() {}

//TODO: Parse account data into total portoflio balance
func (c *CoinbaseClient) GetBalance() ([]byte, error) {
	data, err := c.GetAccounts()
	if err != nil {
		return nil, err
	}
	return data, nil

}
func (c *CoinbaseClient) GetAccounts() ([]byte, error) {

	info := NewRequestInfo{Method: "GET", Uri: "/v2/accounts?&limit=100&order=asc"}
	data, err := c.Get(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

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
