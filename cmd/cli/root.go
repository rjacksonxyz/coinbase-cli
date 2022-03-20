package cli

import (
	client "coinbase-cli/internal"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:  "cbcli",
		Long: "Coinbase CLI: Command-line interface to interact with Coinbase API\nMore info available at: https://github.com/0xmercurial/coinbase-cli/docs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute()
		},
	}
)

func init() {
	//parse command line arguments
	//rootCmd.Flags().StringVarP(&filepath, "ticker", "t", "BTC", "ticker(s) for cryptocurrencies listed on Coinbase")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func execute() error {
	//implement simulation
	return client.CoinbaseCLIRequest()
}
