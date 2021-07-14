package model

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

type ClusterInfo struct {
	ClusterID string
}

func SetCache(key string, value interface{}) {
	Cache.Set(key, value, cache.NoExpiration)
}

func GetCache(key string) (ClusterInfo, bool) {
	var value ClusterInfo
	var found bool

	data, found := Cache.Get(key)
	if found {
		value = data.(ClusterInfo)
	} else {
		fmt.Println("False")
	}
	return value, found
}
