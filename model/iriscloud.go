package model

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

//GetPodsRequest contains Request Parameter
type GetPodsRequest struct {
	Cluster   string
	Namespace string
}

//GetClusterRequest contains Request Parameter
type GetClusterRequest struct {
	Page     int
	Per_page int
}

func GetClusterRequestFromReader(reader io.Reader) (*GetClusterRequest, error) {
	var ret GetClusterRequest
	err := json.NewDecoder(reader).Decode(&ret)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "Failed to decode get clusterID request")
	}

	return &ret, nil
}

func GetPodsRequestFromReader(reader io.Reader) (*GetPodsRequest, error) {
	var ret GetPodsRequest
	err := json.NewDecoder(reader).Decode(&ret)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "Failed to decode get Pods Resource request")
	}
	return &ret, nil
}
