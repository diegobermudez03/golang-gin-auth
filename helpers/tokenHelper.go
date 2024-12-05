package helper

import (
	"fmt"
	"os"
	"strings"
	"time"

	"log"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string 

func init(){
	secretKey = os.Getenv("SECRET_JWT")
	if secretKey == "" {
		secretKey = "secreto"
	}
}

type CustomClaims struct{
	Id			int		`json:"id"`
	Email 		string 	`json:"email"`
	FirstName 	string 	`json:"firstName"`
	LastName 	string	`json:"lastName"`
	Role 		string 	`json:"role"`
	jwt.RegisteredClaims
}

func GenerateAllTokens(email string, id int, firstName string, lastName string, role string) (stringToken string, stringRefreshToken string, err error) {
	tokenClaims := CustomClaims{
		Id: id,
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	stringToken, err = token.SignedString([]byte(secretKey))

	if err != nil{
		return "", "", err
	}
	
	refreshTokenClaims := CustomClaims{
		Id: id,
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*100)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	stringRefreshToken, err = refreshToken.SignedString([]byte(secretKey))

	return stringToken, stringRefreshToken, err
}


func ValidateToken(tokenStr string) bool {
	claims := &CustomClaims{}
	tokenStr = strings.Split(tokenStr," ")[1]

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token)(interface{}, error){
		return []byte(secretKey), nil
	})

	if err != nil{
		log.Println(err)
		return false
	}

	if !token.Valid{

		return false
	}

	if claims.ExpiresAt.Time.Before(time.Now()){
		fmt.Println("Expired token")
	}
	return true
}