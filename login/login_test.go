package login

import (
	"log"
	"testing"
)

func TestClientFromJSON(t *testing.T) {
	client := ClientFromJSON("test-credentials.json")
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal("Unable to get balance: ", err)
	}
	log.Println(balance)
}
