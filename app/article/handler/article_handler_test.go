package handler_test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fjasper13/endpoint-article/app/article/handler"
	"github.com/fjasper13/endpoint-article/app/article/service"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// Failed (NotFinish Yet)
func TestShowArticleHandler(t *testing.T) {
	db, _ := NewMock()
	repo := service.NewArticleRepository(db)
	service := service.NewArticleService(repo)
	handler := handler.NewArticleHandler(service, nil)
	defer db.Close()

	req, err := http.NewRequest("GET", "http://localhost:8000/articles/1", nil)
	fmt.Println("GEGE")

	res := httptest.NewRecorder()
	handler.ShowArticle(res, req)

	assert.NotNil(t, res)
	assert.NoError(t, err)
}
