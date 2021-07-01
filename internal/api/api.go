package api

import "github.com/gorilla/mux"

func Register(rootRouter *mux.Router, context *Context) {
	apiRouter := rootRouter.PathPrefix("/api").Subrouter()
	initIrisCloud(apiRouter, context)
}
