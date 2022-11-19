package controllers

import (
	"encoding/json"
	"example/books-api/connection"
	"example/books-api/dtos"
	"example/books-api/models"
	"fmt"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertUser(user)

	res := dtos.Response{
		ID:      insertID,
		Message: "User registered successfully",
	}

	json.NewEncoder(w).Encode(res)
}
func Login(w http.ResponseWriter, r *http.Request)  {}
func Logout(w http.ResponseWriter, r *http.Request) {}

func insertUser(user models.User) int64 {
	db := connection.CreateConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO users(name,email,password) VALUES($1,$2,$3) RETURNING id"
	var id int64
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("inserted a single record with id %v", id)

	return id
}
