package service_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/service"
	"github.com/stretchr/testify/assert"
)

func TestShowArticleService(t *testing.T) {
	db, mock := NewMock()
	repo := service.NewArticleRepository(db)
	service := service.NewArticleService(repo)
	defer db.Close()

	query := "SELECT id, author, title, body, created_at FROM articles WHERE id = \\?"

	article := entities.Article{
		ID:        22,
		Author:    "TestAuthor",
		Title:     "TestTitle",
		Body:      "TestBody",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).
		AddRow(article.ID, article.Author, article.Title, article.Body, article.CreatedAt)

	mock.ExpectQuery(query).WithArgs(article.ID).WillReturnRows(rows)

	user, err := service.ShowArticle(int(article.ID))
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
