package main

import (
	"fmt"
	"net/http"

	"github.com/rajithacharith/go-rest-template/config"
	"github.com/rajithacharith/go-rest-template/router"
	"github.com/sirupsen/logrus"
)

func main() {
	r, err := router.LoadApiRoutes()
	if err != nil {
		panic(err)
	}
	configs := config.LoadConfigs()
	logrus.Infof("Server is running on port %s", configs.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.Port), r)
}
