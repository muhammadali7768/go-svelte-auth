package router

import (
	"example/books-api/controllers"

	"example/books-api/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//Auth routes
	router.HandleFunc("/api/login", controllers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/register", controllers.Register).Methods("POST", "OPTIONS")

	router.Handle("/api/logout", middlewares.AuthMiddleware(controllers.Logout)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/auth", controllers.IsAuthenticated).Methods("GET", "OPTIONS")
	//Book Routes
	router.Handle("/api/books", middlewares.AuthMiddleware(controllers.CreateBook)).Methods("POST", "OPTIONS")
	router.Handle("/api/books/{id}", middlewares.AuthMiddleware(controllers.GetBookById)).Methods("GET", "OPTIONS")
	router.Handle("/api/books", middlewares.AuthMiddleware(controllers.GetAllBooks)).Methods("GET", "OPTIONS")
	router.Handle("/api/books/{id}", middlewares.AuthMiddleware(controllers.UpdateBook)).Methods("PUT", "OPTIONS")
	router.Handle("/api/books/{id}", middlewares.AuthMiddleware(controllers.DeleteBook)).Methods("DELETE", "OPTIONS")

	return router
}
