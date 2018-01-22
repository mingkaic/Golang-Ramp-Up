package main

import (
	"fmt"
	"net/http"
	"os"
	"./api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/contacts", api.ContactsHandleFunc)
	http.HandleFunc("/api/contacts/", api.ContactHandleFunc)
	fmt.Println("Contacts Web Server started ...")
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8090"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Contacts Application.")
}
