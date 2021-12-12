package cache

import "github.com/fjasper13/endpoint-article/app/article/entities"

type ArticleCache interface {
	Set(key string, value *entities.Article)
	Get(key string) *entities.Article
}
