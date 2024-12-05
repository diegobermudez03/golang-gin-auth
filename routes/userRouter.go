package routes

import (
	controller "github.com/diegobermudez03/golang-jwt-auth/controllers"
	"github.com/diegobermudez03/golang-jwt-auth/middlewares"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router{
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)
	r.Get("/", controller.GetUsers)
	r.Get("/{user_id}", controller.GetUser)
	return r
}