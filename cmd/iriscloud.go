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

		// client := model.NewClient("http://192.168.102.114:32080") //IRIS-CLOUD IP address
		client := model.NewClient("http://localhost:8080") // 내부 서버 IP

		clusterReq := model.GetClusterRequest{
			Page:     0,
			Per_page: -1,
		}
		clusterID, err := client.GetCluster(&clusterReq)

		// podReq := model.GetPodsRequest{ //IRIS-CLOUD request
		// 	Cluster: clusterID,
		// 	//Namespace: os.Getenv("POD_NAMESPACE"),
		// 	Namespace: "iris-cloud",
		// }

		// result, err := client.GetPods(&podReq)
		if err != nil {
			fmt.Printf("Failed to Get Pods from Iris-cloud. ERROR: %s", err.Error())
		}
		fmt.Println(clusterID)
	},
}
