package controller

import (
	"net/http"

	"github.com/rajithacharith/go-rest-template/service"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	service := service.NewHelloWorldService()
	helloWorld := service.GetHelloWorld(ctx)
	w.Write([]byte(helloWorld))
}
