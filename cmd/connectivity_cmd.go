package cmd

import "github.com/spf13/cobra"

var (
	assets string

	connectivityCmd = &cobra.Command{
		Use:   "connectivity",
		Short: "Copy connectivity data to azure",
		Run:   connectivity,
	}
)

func connectivity(cmd *cobra.Command, args []string) {
	println("Hello")
	println("parallel", parallel, assets)
}

func init() {
	connectivityCmd.PersistentFlags().StringVar(&assets, "assets", "", "Csv file containing assets")
}
