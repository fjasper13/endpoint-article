package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fjasper13/endpoint-article/app/article/cache"
	"github.com/fjasper13/endpoint-article/app/article/handler"
	"github.com/fjasper13/endpoint-article/app/article/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file!")
	}

	//Get Connection to Database
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	mySqlUrl := dbUser + ":" + dbPass + "@(127.0.0.1:3306" + ")/" + dbName + "?parseTime=true"

	db, err := sql.Open(dbDriver, mySqlUrl)
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	defer db.Close()

	// Handle Redis
	rdsHost := os.Getenv("REDIS_HOST")
	rdsPassword := os.Getenv("REDIS_PASSWORD")
	rdsExp, err := strconv.Atoi(os.Getenv("REDIS_EXPIRED"))
	if err != nil {
		log.Fatal("Error loading .env Redis file")
	}
	rds := cache.NewRedisCache(rdsHost, 0, rdsPassword, time.Duration(rdsExp)*time.Minute)
	if err != nil {
		log.Fatal("Redis Connection Error")
	}

	newArticleRepo := service.NewArticleRepository(db)
	newArticleService := service.NewArticleService(newArticleRepo)
	newArticleHandler := handler.NewArticleHandler(newArticleService, rds)

	// Handle Router
	router := mux.NewRouter()
	// Handle CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Router v1
	router.HandleFunc("/articles", newArticleHandler.GetArticles).Methods("GET")
	router.HandleFunc("/articles", newArticleHandler.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{article_id}", newArticleHandler.ShowArticle).Methods("GET")

	// =================== LISTENING TO PORT ===================
	fmt.Println("Listen to Port :" + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
