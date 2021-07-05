package cmd

import (
	"fmt"

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

		client := model.NewClient("http://localhost:8080")

		//IRIS-CLOUD /api/cluster
		clusterReq := model.GetClusterRequest{
			Page:     0,
			Per_page: -1,
		}
		clusterID, err := client.GetCluster(&clusterReq)
		if err != nil {
			fmt.Printf("Failed to get CLUSTERID. ERROR: %s\n", err.Error())
		}
		podReq := model.GetPodsRequest{
			Cluster: clusterID,
			//Namespace: os.Getenv("POD_NAMESPACE"),
			Namespace: "iris-cloud",
		}

		result, err := client.GetPods(&podReq)
		if err != nil {
			fmt.Printf("Failed to Get Pods from Iris-cloud. ERROR: %s\n", err.Error())
		}
		fmt.Println(result)
		//fmt.Printf("ClusterID: %s\n Pods Status: %s\n", clusterID, result)
	},
}
