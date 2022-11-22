package controllers

import (
	"encoding/json"
	"errors"
	"example/books-api/connection"
	"example/books-api/dtos"
	"example/books-api/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertUser(user)

	msg := fmt.Sprintf("User registered successfully. User id is %v", insertID)
	res := dtos.Response{
		Status:  http.StatusOK,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}
	regUser, err := findUser(user.Email, user.Password)

	if err != nil {
		res := dtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid credentials! user not found",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	session, _ := store.Get(r, "session-id")
	session.Values["authenticated"] = true
	session.Values["user_id"] = regUser.UserId
	fmt.Println(session)
	// Saves all sessions used during the current request
	session.Save(r, w)

	json.NewEncoder(w).Encode(regUser)

}
func IsAuthenticated(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")

	authenticated := session.Values["authenticated"]

	if authenticated != nil && authenticated != false {
		if userId, exist := session.Values["user_id"]; exist {
			fmt.Printf("User ID %v", userId)
			user, err := findUserById(userId.(int64))
			if err != nil {
				http.Error(w, "User Not Found", http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(user)
			return
		}
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request, user *models.User) {
	session, _ := store.Get(r, "session-id")
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
	res := dtos.Response{
		Status:  http.StatusOK,
		Message: "Logout successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetAuthenticatedUser(r *http.Request) (*models.User, error) {
	//validate the session token in the request,
	session, _ := store.Get(r, "session-id")

	authenticated := session.Values["authenticated"]

	if authenticated != nil && authenticated != false {
		if userId, exist := session.Values["user_id"]; exist {
			fmt.Printf("User ID %v", userId)
			user, err := findUserById(userId.(int64))
			return &user, err
		}
		return nil, errors.New("session not found for this user")
	} else {

		return nil, errors.New("forbidden resources")

	}
}

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

func findUser(email string, password string) (models.User, error) {
	db := connection.CreateConnection()

	defer db.Close()

	sqlStatement := `SELECT id,name,email FROM users where email=$1 AND password=$2 LIMIT 1`

	row := db.QueryRow(sqlStatement, email, password)
	var user models.User
	err := row.Scan(&user.UserId, &user.Name, &user.Email)

	return user, err
}

func findUserById(id int64) (models.User, error) {
	db := connection.CreateConnection()

	defer db.Close()

	sqlStatement := `SELECT id,name,email FROM users where id=$1  LIMIT 1`

	row := db.QueryRow(sqlStatement, id)
	var user models.User
	err := row.Scan(&user.UserId, &user.Name, &user.Email)

	return user, err
}
