package cmd

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/hyobins/service-discovery/internal/api"
	"github.com/hyobins/service-discovery/internal/iriscloud"
	"github.com/hyobins/service-discovery/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Logger *logrus.Logger

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Running Service-Discovery Server",
	Run: func(cmd *cobra.Command, args []string) {
		config := model.GetConfig()

		//Logger = sdlog.GetLogInstance()
		logrus.SetLevel(logrus.DebugLevel)
		logrus.WithFields(logrus.Fields{
			"period":    config.Period,
			"namespace": config.Namespace,
			"brickURL":  config.BrickURL,
		}).Info("Configuratiuon Information")

		r := mux.NewRouter()
		api.Register(r, &api.Context{})

		go func() {
			s := &http.Server{
				Addr:           ":8080",
				Handler:        r,
				ReadTimeout:    10 * time.Second,
				WriteTimeout:   10 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}
			logrus.WithField("addr", s.Addr).Info("Listening")
			s.ListenAndServe()
		}()

		//c := cache.New(5*time.Minute, 10*time.Minute)
		go func() {
			clusterReq := model.GetClusterRequest{
				Page:     0,
				Per_page: -1,
			}
			for {
				data, _ := iriscloud.GetClusterID(&clusterReq)

				clusterinfo := model.ClusterInfo{
					ClusterID: data,
				}
				model.SetCache("data", clusterinfo)
				//c.Set("data", data, cache.DefaultExpiration)
				logrus.Info("Cached ClusterInfo")
				time.Sleep(3 * time.Second)
			}
		}()

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)

		<-ch
		logrus.Info("Shutting Down")

	},
}
