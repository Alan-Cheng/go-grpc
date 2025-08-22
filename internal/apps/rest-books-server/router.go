package restbooksserver

import (
	"go-grpc/internal/pkg/service"

	"github.com/gorilla/mux"
)

func ProvideRouter(bookService service.BookService) *mux.Router {
	r := mux.NewRouter()

	booksHandler := GetNewBookshandler(bookService)

	r.HandleFunc("/books", booksHandler.GetBookList).Methods("GET")
	r.HandleFunc("/books/{isbn:[0-9]+}", booksHandler.GetOrRemoveBookHandler).Methods("GET", "DELETE")
	r.HandleFunc("/books", booksHandler.UpsertBookHandler).Methods("PUT", "POST")

	return r
}
