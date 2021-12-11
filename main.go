package main

import (
	"fmt"
	"net/http"

	"github.com/fjasper13/endpoint-article/app/article/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/articles", handler.GetPosts).Methods("GET")
	router.HandleFunc("/articles", handler.CreatePost).Methods("POST")

	fmt.Println("Listen to Port :8000")
	http.ListenAndServe(":8000", router)
}
