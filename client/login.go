package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login() CoinbaseClient {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("coinbase login - `file` or `manual`:  ")
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimRight(text, "\n"))

	c, err := chooseLoginMethod(text)
	for err != nil {
		//Was looping over the same unchanged text prior
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.TrimRight(text, "\n"))
		c, err = chooseLoginMethod(text)
		if err != nil {
			continue
		}
	}
	return c
}

func chooseLoginMethod(input string) (CoinbaseClient, error) {

	switch input {
	case "file":
		c, err := fileLogin()
		for err != nil {
			c, err = fileLogin()
			if err != nil {
				continue
			}
		}
		return c, nil
	case "manual":
		return ClientFromStdIn(), nil
	default:
		fmt.Print("invalid input - please select `file` or `manual`: ")
		return CoinbaseClient{}, fmt.Errorf("invalid input")
	}
}

func fileLogin() (CoinbaseClient, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter the file path to your Coinbase API credentials:  ")
	filepath, _ := reader.ReadString('\n')
	filepath = strings.TrimRight(filepath, "\n")
	c, err := ClientFromJSON(filepath)
	if err != nil {
		return CoinbaseClient{}, err
	}
	return c, nil
}
