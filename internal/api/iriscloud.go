package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyobins/service-discovery/model"
)

func initIrisCloud(apiRouter *mux.Router, context *Context) {
	addContext := func(handler contextHandlerFunc) *contextHandler {
		return newContextHandler(context, handler)
	}

	iriscloudRouter := apiRouter.PathPrefix("iriscloud").Subrouter()
	iriscloudRouter.Handle("/get", addContext(handleGetIrisCloud)).Methods("GET")
}

func handleGetIrisCloud(c *Context, w http.ResponseWriter, r *http.Request) {
	request, err := model.GetClusterRequestFromReader(r.Body)
	if err != nil {
		c.Logger.WithError(err).Error("Failed to decode get iriscloud cluster ID")
	}

	result, _ := c.Store.GetCluster(request)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	errs := encoder.Encode(result)
	if errs != nil {
		c.Logger.WithError(err).Error("failed to decode result")
	}

}
