package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully with pg")
	return db
}
