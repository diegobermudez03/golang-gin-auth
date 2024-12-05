package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func init(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	conString := os.Getenv("POSTGRES_URL") 
	Db, err = sql.Open("postgres", conString)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to Postgresql")
}
