package service

import (
	"errors"
	"time"

	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/repository"
	"github.com/fjasper13/endpoint-article/app/article/request"
)

func PrepareStoreArticle(req *entities.Article) (res *entities.Article, err error) {
	// Null Handling
	err = ValidateArticle(req)
	if err != nil {
		return nil, err
	}
	req.CreatedAt = time.Now()

	res, err = StoreArticle(req)
	return
}

func StoreArticle(req *entities.Article) (*entities.Article, error) {
	return repository.StoreArticle(req)
}

func ValidateArticle(req *entities.Article) error {
	if req.Author == "" {
		return errors.New("Author can not empty")
	} else if req.Title == "" {
		return errors.New("Title can not empty")

	} else if req.Body == "" {
		return errors.New("Body can not empty")
	}
	return nil
}

func IndexArticle(pr *request.PageRequestStruct) ([]*entities.Article, int, error) {
	return repository.IndexArticle(pr)
}
