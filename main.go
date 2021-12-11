package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fjasper13/endpoint-article/app/article/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file!")
	}

	router := mux.NewRouter()
	// Handle CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/articles", handler.GetPosts).Methods("GET")
	router.HandleFunc("/articles", handler.CreatePost).Methods("POST")

	fmt.Println("Listen to Port :" + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
