package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fjasper13/endpoint-article/app/article/api"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
	"github.com/fjasper13/endpoint-article/app/article/service"
)

type articleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(srv service.ArticleService) *articleHandler {
	return &articleHandler{srv}
}

func (h *articleHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
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

func (h *articleHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
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
