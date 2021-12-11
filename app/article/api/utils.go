package api

import (
	"encoding/json"
	"net/http"
)

func PrintError(err error, w http.ResponseWriter) {
	ret := SingleResponse{}

	switch err.Error() {
	case "400":
		ret.Meta.StatusCode = http.StatusBadRequest
		ret.Meta.Message = []string{"Bad Request"}
		break
	case "401":
		ret.Meta.StatusCode = http.StatusUnauthorized
		ret.Meta.Message = []string{"JWT Token Is Required"}
		break
	case "405":
		ret.Meta.StatusCode = http.StatusMethodNotAllowed
		ret.Meta.Message = []string{"Method Not Allowed"}
		break
	case "500":
		ret.Meta.StatusCode = http.StatusInternalServerError
		ret.Meta.Message = []string{err.Error()}
		break
	default:
		ret.Meta.StatusCode = http.StatusBadRequest
		ret.Meta.Message = []string{err.Error()}
		break
	}

	w.WriteHeader(ret.Meta.StatusCode)
	json.NewEncoder(w).Encode(ret)

	return
}

type SingleResponse struct {
	Data interface{} `json:"data"`
	Meta SimpleMeta  `json:"meta"`
}

type SimpleMeta struct {
	StatusCode int      `json:"status_code"`
	Message    []string `json:"message"`
}

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
