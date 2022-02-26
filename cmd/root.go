package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Crypto exchange service",
	Short: "The crypto exchange application",
	Long:  "Run the 'serve' subcommand to start the http server",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
