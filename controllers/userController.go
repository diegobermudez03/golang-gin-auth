package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/diegobermudez03/golang-jwt-auth/database"
	helper "github.com/diegobermudez03/golang-jwt-auth/helpers"
	"github.com/diegobermudez03/golang-jwt-auth/models"
	"github.com/go-chi/chi/v5"
)

func Signup(w http.ResponseWriter, r *http.Request){
	var user models.User = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return 
	}

	if user.UserType != "USER" && user.UserType != "ADMIN"{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid role"})
		return 
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	//check if the email is already used
	var number int
	_ = database.Db.QueryRowContext(ctx, "SELECT count(*) FROM users WHERE email = $1",  user.Email).Scan(&number)
	if number > 0{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "The email is already registered"})
		return 
	}
	fmt.Println("No duplicate found, continue")

	token, refreshToken, err := helper.GenerateAllTokens(user.Email, user.ID, user.FirstName, user.LastName, user.UserType)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return 
	}

	_, err = database.Db.ExecContext(ctx, "INSERT INTO users(first_name, last_name, password, email, phone, token, user_type, refresh_token) VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		user.FirstName, user.LastName, user.Password, user.Email, user.Phone, token, user.UserType, refreshToken,
	)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return 
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token" : token,
		"refreshRoken" : refreshToken,
	})
}

func Login(w http.ResponseWriter, r *http.Request){

}


func GetUsers(w http.ResponseWriter, r *http.Request){

}

func GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	
	ctx, cancel := context.WithTimeout(r.Context(), 100* time.Second)
	defer cancel()

	var user models.User

	row := database.Db.QueryRowContext(ctx, "SELECT * FROM users WHERE ID = $1 LIMIT 1", userId)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, 
		&user.Phone, &user.Token, &user.UserType, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)
	
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
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
		

}