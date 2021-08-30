package daemon

import (
	"bufio"
	"coinbase-cli/client"
	"fmt"
	"os"
	"strings"
)

type Daemon struct {
	CBClient *client.CoinbaseClient
}

func (d *Daemon) StartCli() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		d.checkInput(text)
	}
}

func (d *Daemon) checkInput(input string) {

	switch strings.TrimRight(input, "\n") {
	case "exit":
		fmt.Println("exiting...")
		os.Exit(1)
	case "accounts":
		accts, err := d.CBClient.GetAccounts()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(accts))

	}
}
