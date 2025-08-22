package restbooksserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-grpc/internal/pkg/model"
	"go-grpc/internal/pkg/service"

	"github.com/gorilla/mux"
)

const (
	SuccessResponseFieldKey = "status"
	ErrorReponseFieldKey    = "error"
)

type BooksHandler struct {
	bookService service.BookService
}

func GetNewBookshandler(bookService service.BookService) *BooksHandler {
	return &BooksHandler{bookService: bookService}
}

func (bh *BooksHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

func (bh *BooksHandler) GetOrRemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	muxVar := mux.Vars(r)
	isbnStr := muxVar["isbn"]
	isbn, err := strconv.Atoi(isbnStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	switch r.Method {
	case "GET":
		book, err := bh.bookService.GetBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, Book(book))
	case "DELETE":
		bh.bookService.RemoveBook(isbn)
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book removed successfully"})
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (bh *BooksHandler) UpsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	switch r.Method {
	case "PUT":
		bh.bookService.AddBook(DBBook(book))
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book added successfully"})
	case "POST":
		bh.bookService.UpdateBook(DBBook(book))
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book updated successfully"})
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{SuccessResponseFieldKey: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
