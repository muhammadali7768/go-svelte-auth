package controllers

import (
	"database/sql"
	"encoding/json"
	"example/books-api/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"messge,omitempty"`
}

func createConnection() *sql.DB {
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertBook(book)

	res := response{
		ID:      insertID,
		Message: "Book created successfully",
	}

	json.NewEncoder(w).Encode(res)
}
func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := getAllBooks()

	if err != nil {
		log.Fatalf("Unable to get all the books. %v", err)
	}

	json.NewEncoder(w).Encode(books)

}
func GetBookById(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}
	var book models.Book

	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := updateBook(int64(id), book)

	msg := fmt.Sprintf("Book updted successfully. Total rows affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {}

func insertBook(book models.Book) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO books(name,price,publisher) VALUES($1,$2,$3) RETURNING id"
	var id int64
	err := db.QueryRow(sqlStatement, book.Name, book.Price, book.Publisher).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("inserted a single record with id %v", id)

	return id

}

func getAllBooks() ([]models.Book, error) {
	db := createConnection()
	defer db.Close()

	var books []models.Book

	sqlStatement := "SELECT * FROM books"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.BookId, &book.Name, &book.Price, &book.Publisher)
		if err != nil {
			log.Fatalf("Unable to scan the rows. %v", err)
		}
		books = append(books, book)
	}
	return books, err

}

func updateBook(id int64, book models.Book) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := "Update books set name=$2, price=$3,publisher=$4 where id=$1"

	res, err := db.Exec(sqlStatement, id, book.Name, book.Price, book.Publisher)

	if err != nil {
		log.Fatalf("Unable to run the update query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to get the affected rows %v", err)
	}

	fmt.Printf(`Total Affected Rows are : %v`, rowsAffected)

	return rowsAffected
}
