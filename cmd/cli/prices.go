package cli

import (
	"github.com/spf13/cobra"
)

var (
	tickers []string
)
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Returns prices for specified cryptocurrency ticker(s)",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.Flags().StringArrayVarP(&tickers, "tickers", "t", []string{"BTC"}, "ticker(s) for cryptocurrencies listed on Coinbase")
}
