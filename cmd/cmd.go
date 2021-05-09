package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	parallel    int
	logLevel    string
	showVersion bool
	version     string
	commit      string
	CazcopyCmd  = &cobra.Command{
		Use:               "cazcopy",
		Short:             "📦 cazcopy - copy data from cassandra to azure storage",
		Long:              ``,
		SilenceErrors:     true,
		SilenceUsage:      true,
		PersistentPreRunE: readConfig,
		PreRunE:           preFlight,
		RunE:              start,
	}
)

func readConfig(ccmd *cobra.Command, args []string) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

func preFlight(ccmd *cobra.Command, args []string) error {
	if showVersion {
		fmt.Printf("cazcopy %s (%s)\n", version, commit)
		return fmt.Errorf("")
	}

	return nil
}

func start(ccmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	viper.SetDefault("CAZCOPY_LOG_LEVEL", "info")

	CazcopyCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Display the current version of this CLI")
	CazcopyCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "", "log level")
	CazcopyCmd.PersistentFlags().IntVar(&parallel, "parallel", 1, "number parallel tasks")

	CazcopyCmd.AddCommand(connectivityCmd)
}
