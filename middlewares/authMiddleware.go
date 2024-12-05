package middlewares

import (
	"encoding/json"
	"net/http"

	helper "github.com/diegobermudez03/golang-jwt-auth/helpers"
)


func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			token := r.Header.Get("Authorization")
			if !helper.ValidateToken(token) {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error" : "Invalid or expired token",
				})
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}
