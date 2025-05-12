package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactService struct {
	Contacts map[int]Contact
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
