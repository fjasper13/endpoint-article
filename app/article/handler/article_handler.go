package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fjasper13/endpoint-article/app/article/api"
	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/fjasper13/endpoint-article/app/article/request"
	"github.com/fjasper13/endpoint-article/app/article/service"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Get Data From Body Request
	var request *entities.Article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		api.PrintError(err, w)
		return
	}

	// Create New Article
	result, err := service.PrepareStoreArticle(request)
	if err != nil {
		api.PrintError(err, w)
		return
	}

	// Prepare response
	res := api.NewResponse(result)
	api.SendJSON(res, w)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {

	//Convert Page Request
	param := r.URL.Query()
	pr := request.PageRequest(param)
	fetch, _, err := service.IndexArticle(pr)
	if err != nil {
		api.PrintError(err, w)
		return
	}

	// Prepare response
	res := api.NewResponse(fetch)
	api.SendJSON(res, w)

}
