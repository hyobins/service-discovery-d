package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service-discovery",
	Short: "Service Discovery is blah bla",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
		//serverCmd.RunE(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(iriscloudCmd)
	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	rootCmd.Execute()
}
