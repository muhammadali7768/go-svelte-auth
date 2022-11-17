package main

import (
	"example/books-api/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Server on Port 8888")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8888", handler))
}
