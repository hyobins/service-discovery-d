package iriscloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hyobins/service-discovery/model"
	"github.com/pkg/errors"
)

func closeBody(r *http.Response) {
	if r.Body != nil {
		_, _ = ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
	}
}

//GetClusterID returns ClusterID through http request to IRIS-CLOUD
func GetClusterID(request *model.GetClusterRequest) (string, error) {
	resp, err := http.Get("http://192.168.102.114:32080/api/cluster")
	if err != nil {
		fmt.Printf("ERROR with %s", err)
		return "", err
	}
	defer closeBody(resp)

	switch resp.StatusCode {
	case http.StatusOK:

		m := []map[string]interface{}{}
		bytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bytes, &m)
		if err != nil {
			fmt.Print("Unmarshal ERROR: ", err)
		}
		result, _ := json.Marshal(m[0]["id"])
		clusterID := strings.Trim(string(result), "\"")
		return clusterID, nil
	default:
		return "", errors.Errorf("Failed with status code %d", resp.StatusCode)
	}
}

func GetPodResource(request *model.GetPodsRequest) (string, error) {
	req, err := http.NewRequest("GET", "http://192.168.102.114:32080/apiv2/state/pods", nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	result := strings.Trim(request.Cluster, "\n")
	clusterID := strings.Trim(result, "\"")
	q.Add("cluster", clusterID)
	q.Add("namespace", request.Namespace)
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())

	if err != nil {
		return "", err
	}
	defer closeBody(resp)

	switch resp.StatusCode {
	case http.StatusOK:
		bytes, _ := ioutil.ReadAll(resp.Body)
		return string(bytes), nil
	default:
		return "", errors.Errorf("failed with status code %d", resp.StatusCode)
	}
}
