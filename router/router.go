package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rajithacharith/go-rest-template/controller"
)

func LoadApiRoutes() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controller.HelloWorld)

	return r, nil
}
