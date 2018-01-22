package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/snarad/git_test/golangWithUnitTest/books/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/book/", api.BookHandleFunc)
	fmt.Println("Books Web Server started ...")
	http.ListenAndServe(port(), nil)
}

// Check is the PORT key exists in the system environment variables
// if not assign default port 8090
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8090"
	}
	return ":" + port
}

// index page should just print Books Application on the browser
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Books Application.")
}