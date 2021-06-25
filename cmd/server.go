package cmd

import "github.com/spf13/cobra"

func init() {

}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server cmd",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
	},
}
