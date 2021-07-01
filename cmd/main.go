package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service-discovery",
	Short: "Service Discovery is blah bla",
	Run: func(cmd *cobra.Command, args []string) {
		//serverCmd.RunE(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(iriscloudCmd)
}

func Execute() {
	rootCmd.Execute()
}
