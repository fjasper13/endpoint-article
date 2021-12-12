package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/fjasper13/endpoint-article/app/article/api"
	"github.com/fjasper13/endpoint-article/app/article/cache"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
	"github.com/fjasper13/endpoint-article/app/article/service"
)

var (
	articleService service.ArticleService
	articleCache   cache.ArticleCache
)

type articleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(srv service.ArticleService, cache cache.ArticleCache) *articleHandler {
	articleService = srv
	articleCache = cache
	return &articleHandler{srv}
}

func (h *articleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Get Data From Body Request
	var request *entities.Article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		api.PrintError(err, w)
		return
	}

	// Create New Article
	result, err := h.service.PrepareStoreArticle(request)
	if err != nil {
		api.PrintError(err, w)
		return
	}

	// Prepare response
	res := api.NewResponse(result)
	api.SendJSON(res, w)
}

func (h *articleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	//Convert Page Request
	param := r.URL.Query()
	pr := request.PageRequest(param)
	fetch, count, err := h.service.IndexArticle(pr)
	if err != nil {
		api.PrintError(err, w)
		return
	}

	// Prepare response
	res := api.SetResponsePagination(pr, fetch, count)
	api.SendJSON(res, w)

}

func (h *articleHandler) ShowArticle(w http.ResponseWriter, r *http.Request) {
	//Convert Page Request
	articleID := strings.Split(r.URL.Path, "/")[2]
	intID, _ := strconv.Atoi(articleID)

	var cArticle *entities.Article = articleCache.Get(articleID)

	if cArticle == nil {
		fetch, err := h.service.ShowArticle(intID)
		if err != nil {
			api.PrintError(err, w)
			return
		}
		articleCache.Set(articleID, fetch)
		// Prepare response
		res := api.NewResponse(fetch)
		api.SendJSON(res, w)
	} else {
		res := api.NewResponse(cArticle)
		api.SendJSON(res, w)
	}

}
