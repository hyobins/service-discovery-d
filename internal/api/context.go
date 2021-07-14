package api

import (
	"github.com/sirupsen/logrus"
)

//type Store interface {
//	GetClusterID(request *model.GetClusterRequest) (string, error)
//	GetPodResource(request *model.GetPodsRequest) (string, error)
//}

type Context struct {
	//Store  Store
	Logger logrus.FieldLogger
}

func (c *Context) Clone() *Context {
	return &Context{
		//	Store:  c.Store,
		Logger: c.Logger,
	}
}
