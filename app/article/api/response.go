package api

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/fjasper13/endpoint-article/app/article/request"
)

func NewResponse(data interface{}) *SingleResponse {
	response := &SingleResponse{
		Data: data,
		Meta: SimpleMeta{
			StatusCode: http.StatusOK,
			Message:    []string{"Success"},
		},
	}
	return response
}

func SendJSON(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func SetResponsePagination(pageRequest *request.PageRequestStruct, data interface{}, count int) interface{} {
	response := &PaginateResponse{
		Data: data,
		Meta: SimpleMeta{
			StatusCode: 200,
			Message:    []string{"Success"},
		},
		CurrentPage: pageRequest.Page,
		LastPage:    int64(math.Ceil(float64(count) / float64(pageRequest.PerPage))),
		PerPage:     pageRequest.PerPage,
		Total:       int64(count),
	}

	return response
}

type PaginateResponse struct {
	CurrentPage int64       `json:"current_page"`
	LastPage    int64       `json:"last_page"`
	PerPage     int64       `json:"per_page"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
	Meta        SimpleMeta  `json:"meta"`
}
