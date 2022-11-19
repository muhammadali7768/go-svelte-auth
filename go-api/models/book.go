package models

type Book struct {
	BookId    int64   `json:"bookid"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Publisher string  `json:"publisher"`
}
