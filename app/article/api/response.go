package api

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"strconv"

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

func SetPaginate(r *http.Request, pageRequest *request.PageRequestStruct, data interface{}, count int) interface{} {
	var res interface{}

	buffer := new(bytes.Buffer)
	isFirstParam := true
	for k, v := range r.URL.Query() {
		if k != "page" {
			if !isFirstParam {
				buffer.WriteString("&")
			} else {
				isFirstParam = false
			}
			buffer.WriteString(k)
			buffer.WriteString("=")
			buffer.WriteString(v[0])
		}
	}

	if pageRequest.Paginate == 1 {
		res = setResponsePagination(pageRequest, data, count, buffer.String(), isFirstParam)
	} else {
		res = NewResponse(data)
		return res
	}

	return res
}

func setResponsePagination(pageRequest *request.PageRequestStruct, data interface{}, count int, params string, isFirstParam bool) interface{} {
	response := &VueTablePaginateResponse{
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

	if pageRequest.Page >= 1 && pageRequest.Page < response.LastPage {
		nextPage := "page=" + strconv.Itoa(int(pageRequest.Page)+1)
		if !isFirstParam {
			nextPage = "&" + nextPage
		}
		// response.NextPageUrl = env.Get("APP_HOST") + pageUrl + "?" + params + nextPage
	}

	if pageRequest.Page > 1 && pageRequest.Page <= response.LastPage {
		prevPage := "page=" + strconv.Itoa(int(pageRequest.Page)-1)
		if !isFirstParam {
			prevPage = "&" + prevPage
		}
		// response.PrevPageUrl = env.Get("APP_HOST") + pageUrl + "?" + params + prevPage
	}

	return response
}

type VueTablePaginateResponse struct {
	CurrentPage int64       `json:"current_page"`
	LastPage    int64       `json:"last_page"`
	PerPage     int64       `json:"per_page"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
	Meta        SimpleMeta  `json:"meta"`
}
