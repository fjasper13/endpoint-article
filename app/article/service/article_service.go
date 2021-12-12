package service

import (
	"errors"
	"time"

	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
)

type ArticleService interface {
	PrepareStoreArticle(req *entities.Article) (response *entities.Article, err error)
	IndexArticle(pr *request.PageRequestStruct) (res []*entities.Article, count int, err error)
	ShowArticle(ID int) (res *entities.Article, err error)
}

type articleService struct {
	articleRepository ArticleRepository
}

func NewArticleService(repository ArticleRepository) *articleService {
	return &articleService{repository}
}

func (s *articleService) PrepareStoreArticle(req *entities.Article) (res *entities.Article, err error) {
	// Null Handling
	err = ValidateArticle(req)
	if err != nil {
		return nil, err
	}
	req.CreatedAt = time.Now()

	res, err = s.StoreArticle(req)
	return
}

func (s *articleService) StoreArticle(req *entities.Article) (*entities.Article, error) {
	return s.articleRepository.StoreArticle(req)
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

func (s *articleService) IndexArticle(pr *request.PageRequestStruct) ([]*entities.Article, int, error) {
	return s.articleRepository.IndexArticle(pr)
}

func (s *articleService) ShowArticle(ID int) (*entities.Article, error) {
	return s.articleRepository.ShowArticle(ID)
}
