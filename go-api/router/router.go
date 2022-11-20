package router

import (
	"example/books-api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//Auth routes
	router.HandleFunc("/api/login", controllers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/logout", controllers.Logout).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/register", controllers.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth", controllers.IsAuthenticated).Methods("GET", "OPTIONS")
	//Book Routes
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
