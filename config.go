package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	Period    string `yaml: "period"`
	BrickURL  string `yaml: "brick_url"`
	Namespace string `yaml: "namespace"`
	LogLevel  string `yaml: "loglevel"`
}

var config Config

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "<path-to-kubeconfig>")
	if err != nil {
		fmt.Errorf("ERROR: ", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Errorf("ERROR: ", err)
	}
	// 파드를 나열하기 위해 API에 접근한다
	pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func GetConfig() Config {
	return config
}
