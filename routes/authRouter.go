package routes

import (
	controller "github.com/diegobermudez03/golang-jwt-auth/controllers"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() chi.Router{
	r := chi.NewRouter()
	r.Post("/signup", controller.Signup)
	r.Post("/login", controller.Login)
	return r
}