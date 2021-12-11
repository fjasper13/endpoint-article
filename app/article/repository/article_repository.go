package repository

import (
	"errors"

	"github.com/fjasper13/endpoint-article/app/article/database"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
)

func StoreArticle(req *entities.Article) (response *entities.Article, err error) {
	// Get DB Connection
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Insert Statement
	insert, err := db.Prepare("INSERT INTO articles(author,title,body,created_at) VALUES(?,?,?,?)")
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

func IndexArticle(pr *request.PageRequestStruct) (res []*entities.Article, count int, err error) {
	// Get DB Connection
	db, err := database.GetDB()
	if err != nil {
		return nil, 0, err
	}
	sql := "SELECT id, author, title, body, created_at FROM articles"

	// Handle Page Request
	sql += database.BuildQuery(pr)

	// Fetch Data
	fetch, err := db.Query(sql)
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
	sqlCount := "SELECT COUNT(*) FROM articles"
	sqlCount += database.BuildQuery(pr)

	err = db.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return nil, 0, errors.New("500")
	}

	return
}
