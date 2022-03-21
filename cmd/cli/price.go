package cli

import (
	client "coinbase-cli/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	currency_pairs string
)
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Returns prices for specified cryptocurrency ticker(s)",
	Run: func(cmd *cobra.Command, args []string) {
		if c, err := client.NewClient(filepath); err != nil {
			fmt.Println(err)
		} else {
			c.GetPrice(client.PriceRequest{
				CurrencyPairs: currency_pairs,
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.Flags().StringVarP(&currency_pairs, "currency_pairs", "c", "BTC-USD", "comma separated list of currency pairs (for pair A-B, returns amount of A denominated in B)")
}
