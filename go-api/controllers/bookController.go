package controllers

import (
	"database/sql"
	"encoding/json"
	"example/books-api/connection"
	"example/books-api/dtos"
	"example/books-api/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertBook(book)

	msg := fmt.Sprintf("Book created successfully. Book id is %v", insertID)
	res := dtos.Response{
		Status:  http.StatusOK,
		Message: msg,
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
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	book, err := getBookById(int64(id))

	if err != nil {
		log.Fatalf("Unable to get the book by id. %v", err)
	}

	json.NewEncoder(w).Encode(book)
}

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

	res := dtos.Response{
		Status:  http.StatusOK,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deleteBook(int64(id))

	msg := fmt.Sprintf("Book Deleted successfully. Total rows affected %v", deletedRows)

	res := dtos.Response{
		Status:  http.StatusOK,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertBook(book models.Book) int64 {
	db := connection.CreateConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO books(name,price,publisher,user_id) VALUES($1,$2,$3,$4) RETURNING id"
	var id int64
	println("User ID: %v", book.UserId)
	println("User ID: %v", book.Name)
	err := db.QueryRow(sqlStatement, book.Name, book.Price, book.Publisher, book.UserId).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("inserted a single record with id %v", id)

	return id

}

func getAllBooks() ([]models.Book, error) {
	db := connection.CreateConnection()
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

		err = rows.Scan(&book.BookId, &book.Name, &book.Price, &book.Publisher, &book.UserId)
		if err != nil {
			log.Fatalf("Unable to scan the rows. %v", err)
		}
		books = append(books, book)
	}
	return books, err

}

func updateBook(id int64, book models.Book) int64 {
	db := connection.CreateConnection()
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

func getBookById(id int64) (models.Book, error) {
	db := connection.CreateConnection()
	defer db.Close()

	var book models.Book

	sqlStatement := "SELECT * FROM books where id=$1"

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&book.BookId, &book.Name, &book.Price, &book.Publisher)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
		return book, nil
	case nil:
		return book, nil
	default:
		log.Fatalf("Unable to scan the rows. %v", err)
	}

	return book, err
}

func deleteBook(id int64) int64 {

	db := connection.CreateConnection()
	defer db.Close()

	sqlStatement := "DELETE FROM books  where id=$1"

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to run the delete query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to get the affected rows %v", err)
	}

	fmt.Printf("Total Affected Rows are : %v", rowsAffected)

	return rowsAffected
}
