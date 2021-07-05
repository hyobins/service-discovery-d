package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyobins/service-discovery/internal/iriscloud"
	"github.com/hyobins/service-discovery/model"
)

func initIrisCloud(apiRouter *mux.Router, context *Context) {
	addContext := func(handler contextHandlerFunc) *contextHandler {
		return newContextHandler(context, handler)
	}

	iriscloudRouter := apiRouter.PathPrefix("/iriscloud").Subrouter()
	iriscloudRouter.Handle("/get", addContext(handleGetIrisCloud)).Methods("GET")
	iriscloudRouter.Handle("/getpods", addContext(handlerGetPodResource)).Methods("POST")
}

func handleGetIrisCloud(c *Context, w http.ResponseWriter, r *http.Request) {
	request, err := model.GetClusterRequestFromReader(r.Body)
	if err != nil {
		c.Logger.WithError(err).Error("Failed to decode get iriscloud cluster ID")
	}

	result, err := iriscloud.GetClusterID(request)

	if err != nil {
		fmt.Print(err)
		c.Logger.WithError(err).Error("Failed to use  getcluster api")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	errs := encoder.Encode(result)
	if errs != nil {
		c.Logger.WithError(err).Error("failed to decode result")
	}
}

func handlerGetPodResource(c *Context, w http.ResponseWriter, r *http.Request) {
	request, err := model.GetPodsRequestFromReader(r.Body)
	if err != nil {
		c.Logger.WithError(err).Error("Failed to decode get iriscloud pod resource")
	}

	result, err := iriscloud.GetPodResource(request)
	if err != nil {
		fmt.Print(err)
		c.Logger.WithError(err).Error("Failed to use GetPodResource api")
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	errs := encoder.Encode(result)
	if errs != nil {
		c.Logger.WithError(err).Error("Failed to decode result")
	}

	fmt.Printf("Pod Resource: %s\n", result)
}
