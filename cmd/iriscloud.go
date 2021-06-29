package cmd

import (
	"fmt"
	"os"

	"github.com/hyobins/service-discovery/model"

	"github.com/spf13/cobra"
)

func init() {
	//statusGetCmd.Flags().String("range", "current", "status 범위")

	iriscloudCmd.AddCommand(statusGetCmd)
}

var iriscloudCmd = &cobra.Command{
	Use:   "iris-cloud",
	Short: "Get Pod's Status through IRIS-CLOUD",
}

var statusGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get status information",
	Run: func(cmd *cobra.Command, args []string) {
		//getRange, _ := command.Flags().GetString("range")
		fmt.Println("get!! ")

		client := model.NewClient("addr") //IRIS-CLOUD IP address

		clusterReq := model.GetClusterRequest{
			Page:     0,
			Per_page: -1,
		}
		clusterID, err := client.GetCluster(&clusterReq)

		podReq := model.GetPodsRequest{ //IRIS-CLOUD request
			Cluster:   clusterID,
			Namespace: os.Getenv("POD_NAMESPACE"),
		}

		result, err := client.GetPods(&podReq)
		if err != nil {
			fmt.Printf("Failed to Get Pods from Iris-cloud. ERROR: %s", err.Error())
		}
		fmt.Printf(result)
	},
}
