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

type GetClusterRequest struct {
	Page     int
	Per_page int
}

func GetClusterRequestFromReader(reader io.Reader) (*GetClusterRequest, error) {
	var ret GetClusterRequest
	err := json.NewDecoder(reader).Decode(&ret)
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "failed to decode get clusterID request")
	}

	return &ret, nil
}
