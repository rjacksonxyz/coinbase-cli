package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TODO: Finish Login
func Login() CoinbaseClient {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("coinbase login - `file` or `manual`?:  ")
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimRight(text, "\n"))

	c, err := chooseLoginMethod(text)
	for err != nil {
		c, err = chooseLoginMethod(text)
		if err != nil {
			continue
		}
		err = nil
	}
	return c
}

func chooseLoginMethod(input string) (CoinbaseClient, error) {

	switch input {
	case "file":
		return fileLogin(), nil

	case "manual":
		return ClientFromStdIn(), nil
	default:
		fmt.Println("invalid input - please select `file` or `manual`")
		return CoinbaseClient{}, fmt.Errorf("invalid input")
	}
}

func fileLogin() CoinbaseClient {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter the file path to your Coinbase API credentials:  ")
	filepath, _ := reader.ReadString('\n')
	filepath = strings.TrimRight(filepath, "\n")
	c := ClientFromJSON(filepath)
	return c
}
