package model

//GetPodsRequest contains Request Parameter
type GetPodsRequest struct {
	Cluster   string
	Namespace string
}

type GetClusterRequest struct {
	Page     int
	Per_page int
}
