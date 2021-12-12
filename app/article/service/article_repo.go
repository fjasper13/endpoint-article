package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/fjasper13/endpoint-article/app/article/database"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
)

type ArticleRepository interface {
	StoreArticle(req *entities.Article) (response *entities.Article, err error)
	IndexArticle(pr *request.PageRequestStruct, sql string, sqlCount string) (res []*entities.Article, count int, err error)
	ShowArticle(ID int) (res *entities.Article, err error)
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) StoreArticle(req *entities.Article) (response *entities.Article, err error) {
	// Insert Statement
	insert, err := r.db.Prepare("INSERT INTO articles(author,title,body,created_at) VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	result, err := insert.Exec(req.Author, req.Title, req.Body, req.CreatedAt)
	if err != nil {
		return nil, err
	}

	req.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	response = req

	return
}

func (r *articleRepository) IndexArticle(pr *request.PageRequestStruct, sql string, sqlCount string) (res []*entities.Article, count int, err error) {
	// Handle Page Request
	sql += database.BuildQuery(pr)
	// Fetch Data
	fetch, err := r.db.Query(sql)
	if err != nil {
		return nil, 0, err
	}

	defer fetch.Close()

	res = []*entities.Article{}

	// Decode and Append to Array
	for fetch.Next() {
		article := entities.Article{}
		err = fetch.Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.CreatedAt)
		if err != nil {
			return nil, 0, errors.New("500")
		}
		res = append(res, &article)
	}

	// Count All Articles
	sqlCount += database.BuildQuery(pr)

	err = r.db.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return nil, 0, errors.New("500")
	}

	return
}

func (r *articleRepository) ShowArticle(ID int) (res *entities.Article, err error) {
	article := entities.Article{}
	// Fetch Data
	err = r.db.QueryRow("SELECT id, author, title, body, created_at FROM articles WHERE id = ? AND deleted_at IS NULL", ID).Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No Article on ID : %d", ID))
		}
		return nil, err
	}
	res = &article

	return
}
