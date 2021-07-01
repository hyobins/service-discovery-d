package iriscloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyobins/service-discovery/model"
	"github.com/pkg/errors"
)

func closeBody(r *http.Response) {
	if r.Body != nil {
		_, _ = ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
	}
}

//iriscloud에 접속하여 받아오기,,
func GetClusterI(request *model.GetClusterRequest) (string, error) {
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
		clusterID, _ := json.Marshal(m[0]["id"])
		return string(clusterID), nil
	default:
		return "", errors.Errorf("failed with status code %d", resp.StatusCode)
	}
}
