package main

import (
	"bufio"
	cl "coinbase-cli/client"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkInput(input string) {

	switch strings.TrimRight(input, "\n") {
	case "exit":
		log.Print("exiting...")
		os.Exit(1)
	}
}

func main() {
	c := cl.Login()
	fmt.Print(c)
	// CLI daemon
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		checkInput(text)
	}
}

/* Outline Functionality

- Finish Login setup
- Show overall account balance, refreshed every 5 secs
- Show account balances, broken down by asset allocation
- Show individual assets
- Show multiple live feeds

*/
