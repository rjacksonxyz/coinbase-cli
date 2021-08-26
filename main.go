package main

import (
	"bufio"
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
	// CLI daemon
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		checkInput(text)
	}
}

/* Outline Functionality

- Show account balances, broken down by asset allocation
- Show individual assets

*/
