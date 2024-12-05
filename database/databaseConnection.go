package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
		return
	}
	fmt.Println("ENV variables loaded")

	conString := os.Getenv("POSTGRES_URL") 
	fmt.Printf("Connecting to postgres on %s\n", conString)
	
	Db, err = sql.Open("postgres", conString)

	if err = Db.Ping(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to Postgresql")
}
