package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/diegobermudez03/golang-jwt-auth/database"
	"github.com/diegobermudez03/golang-jwt-auth/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	defer database.Db.Close()
	defer fmt.Println("closing database")

	port := os.Getenv("PORT")

	if port == ""{
		port= "8000"
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: 	[]string{"*"},
		AllowedMethods:		[]string{"GET", "POST", "PUT", "DELETE"},
	}))

	//mounting routes
	router.Mount("/users/auth", routes.AuthRoutes())
	router.Mount("/users", routes.UserRoutes())

	//health check
	router.Get("/api-1", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"success": "Access granted for api 1",
		})
	})

	router.Get("/api-2", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"success": "Access granted for api 2",
		})
	})

	srv := &http.Server{
		Handler: router,
		Addr : ":" + port,
	}

	err := srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}