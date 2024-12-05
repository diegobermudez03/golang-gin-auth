package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")

	if port == ""{
		port= "8000"
	}

	router := chi.NewRouter()

	//

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