package service_test

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
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

func TestShowArticle(t *testing.T) {
	db, mock := NewMock()
	repo := service.NewArticleRepository(db)
	defer db.Close()

	query := "SELECT id, author, title, body, created_at FROM articles WHERE id = \\?"

	article := entities.Article{
		ID:        13,
		Author:    "TestAuthor",
		Title:     "TestTitle",
		Body:      "TestBody",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).
		AddRow(article.ID, article.Author, article.Title, article.Body, article.CreatedAt)

	mock.ExpectQuery(query).WithArgs(article.ID).WillReturnRows(rows)

	user, err := repo.ShowArticle(int(article.ID))
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

// Failed (NotFinish Yet)
func TestIndexArticle(t *testing.T) {
	db, mock := NewMock()
	repo := service.NewArticleRepository(db)
	defer db.Close()

	query := "SELECT id, author, title, body, created_at FROM articles"

	article := entities.Article{
		ID:        2,
		Author:    "TestAuthor",
		Title:     "TestTitle",
		Body:      "TestBody",
		CreatedAt: time.Now(),
	}

	pr := request.PageRequestStruct{
		Page:     1,
		Paginate: 1,
		PerPage:  3,
	}

	rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).
		AddRow(article.ID, article.Author, article.Title, article.Body, article.CreatedAt)

	mock.ExpectQuery(query).WithArgs(pr).WillReturnRows(rows)

	user, _, err := repo.IndexArticle(&pr)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

// Failed (NotFinish Yet)
func TestStoreArticle(t *testing.T) {
	db, mock := NewMock()
	repo := service.NewArticleRepository(db)
	defer db.Close()

	query := "INSERT INTO articles (author,title,body,created_at) VALUES (?,?,?,?)"

	article := entities.Article{
		ID:        13,
		Author:    "FF",
		Title:     "QQ",
		Body:      "PP",
		CreatedAt: time.Now(),
	}

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(article.Author, article.Title, article.Body, article.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := repo.StoreArticle(&article)
	assert.NoError(t, err)
}
