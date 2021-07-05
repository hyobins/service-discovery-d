package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	address    string
	httpClient *http.Client
}

func NewClient(address string) *Client {
	return &Client{
		address:    address,
		httpClient: &http.Client{},
	}
}

func closeBody(r *http.Response) {
	if r.Body != nil {
		_, _ = ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
	}
}

func (c *Client) doGet(u string) (*http.Response, error) {
	return c.httpClient.Get(u)
}

func (c *Client) doPost(u string, request interface{}) (*http.Response, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Post(u, "application/json", bytes.NewReader(requestBytes))
}

func (c *Client) buildURL(urlPath string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", c.address, fmt.Sprintf(urlPath, args...))
}

//GetCluster 는 외부에서 해당 경로 요청시 내부 서버에서 처리해주는 router
func (c *Client) GetCluster(req *GetClusterRequest) (string, error) {
	resp, err := c.doGet(c.buildURL("/api/iriscloud/get"))
	if err != nil {
		return "", err
	}
	defer closeBody(resp)

	switch resp.StatusCode {
	case http.StatusOK:
		bytes, _ := ioutil.ReadAll(resp.Body)
		return string(bytes), nil
	default:
		return "", errors.Errorf("failed with statuscode %d", resp.StatusCode)
	}
}

func (c *Client) GetPods(req *GetPodsRequest) (string, error) {
	resp, err := c.doPost(c.buildURL("/api/iriscloud/getpods"), req)
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

// func (c *Client) GetCluster(req *GetClusterRequest) (string, error) {
// 	resp, err := c.doGet(c.buildURL("/api/cluster"))
// 	if err != nil {
// 		fmt.Printf("ERROR with %s", err)
// 		return "", err
// 	}
// 	defer closeBody(resp)

// 	switch resp.StatusCode {
// 	case http.StatusOK:

// 		m := []map[string]interface{}{}
// 		bytes, _ := ioutil.ReadAll(resp.Body)
// 		err := json.Unmarshal(bytes, &m)
// 		if err != nil {
// 			fmt.Print("Unmarshal ERROR: ", err)
// 		}
// 		clusterID, _ := json.Marshal(m[0]["id"])
// 		return string(clusterID), nil
// 	default:
// 		return "", errors.Errorf("failed with status code %d", resp.StatusCode)
// 	}
// }
