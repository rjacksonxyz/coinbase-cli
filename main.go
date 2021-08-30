package main

import (
	cl "coinbase-cli/client"
	"coinbase-cli/daemon"
)

func main() {
	c := cl.Login()
	d := daemon.Daemon{CBClient: &c}
	// CLI daemon
	d.StartCli()
}

/* Outline Functionality

- Finish Login setup
- Show overall account balance, refreshed every 5 secs
- Show account balances, broken down by asset allocation
- Show individual assets
- Show multiple live feeds

*/
