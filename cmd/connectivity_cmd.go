package cmd

import (
	"github.com/Azure/azure-storage-blob-go/azblob"
	ct "github.com/bluecolor/connectivity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	assets string

	connectivityCmd = &cobra.Command{
		Use:   "connectivity",
		Short: "Copy connectivity data to azure",
		Run:   connectivity,
	}
)

func getTypeIds() (typeids []int, err error) {
	return typeids, err
}

func connectivity(cmd *cobra.Command, args []string) {
	accountname := viper.GetString("AZURE_STORAGE_ACCOUNT")
	accountkey := viper.GetString("AZURE_STORAGE_ACCESS_KEY")
	credential, err := azblob.NewSharedKeyCredential(accountname, accountkey)
	if err != nil {
		println("Unable to create azure credential")
	}
	c := ct.NewConnectivity("assets.csv", "201801", "202001", typeids)
}

func init() {
	connectivityCmd.PersistentFlags().StringVar(&assets, "assets", "", "Csv file containing assets")
}
