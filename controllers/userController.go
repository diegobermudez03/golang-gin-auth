package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/diegobermudez03/golang-jwt-auth/database"
	helper "github.com/diegobermudez03/golang-jwt-auth/helpers"
	"github.com/diegobermudez03/golang-jwt-auth/models"
	"github.com/go-chi/chi/v5"
)

func Signup(w http.ResponseWriter, r *http.Request){

}

func Login(w http.ResponseWriter, r *http.Request){

}


func GetUsers(w http.ResponseWriter, r *http.Request){

}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	userId := chi.URLParam(r, "user_id")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := helper.MatchUserTypeToUID(user, userId); err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	ctx, cancel := context.WithTimeout(r.Context(), 100* time.Second)
	defer cancel()

	row := database.Db.QueryRowContext(ctx, "SELECT * FROM users WHERE ID = $1 LIMIT 1", userId)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, 
		&user.Email, &user.Phone, &user.Token, &user.UserType, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)
	
	if err != nil{
		if err == sql.ErrNoRows{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "No user found"})
			return 
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return 
	}
	json.NewEncoder(w).Encode(user)
		

}