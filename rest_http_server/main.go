package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	"github.com/rewardStyle/campaign-service/lib/log"
	"github.com/rewardStyle/generic-service/controllers"
	"github.com/rewardStyle/generic-service/services"
	"github.com/spf13/viper"
	"github.com/stvp/rollbar"
)

var app newrelic.Application

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Fatal error config file: %s \n", err))
	}

	if err := log.Init(viper.GetString("environment"), viper.GetString("log")); err != nil {
		log.Fatal(fmt.Sprintf("Fatal error log : %s \n", err))
	}

	rollbar.Token = viper.GetString("rollbar.token")
	rollbar.Environment = viper.GetString("environment")

	nwConfig := newrelic.NewConfig(viper.GetString("newrelic.appname"), viper.GetString("newrelic.licence"))

	var err error
	if app, err = newrelic.NewApplication(nwConfig); err != nil {
		log.Fatal(fmt.Sprintf("Fatal error NewRelic NewApplication: %s \n", err))
	}
}

func main() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()

	repository := services.NewGenericRepository()
	svc := services.NewGenericService(repository)
	ctrl := controllers.NewGenericController(svc)

	v1.HandleFunc("/generics", ctrl.GetGeneric).Methods("GET")

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(fmt.Sprintf("Fatal error server.Serve: %s \n", err))
	}
}
