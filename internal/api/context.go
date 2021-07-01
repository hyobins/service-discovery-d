package api

import (
	"github.com/hyobins/service-discovery/model"
	"github.com/sirupsen/logrus"
)

type Store interface {
	GetCluster(request *model.GetClusterRequest) (string, error)
	GetPods(request *model.GetPodsRequest) (string, error)
}

type Context struct {
	Store  Store
	Logger logrus.FieldLogger
}

func (c *Context) Clone() *Context {
	return &Context{
		Store:  c.Store,
		Logger: c.Logger,
	}
}
