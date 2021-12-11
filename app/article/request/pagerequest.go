package request

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type PageRequestStruct struct {
	Page     int64
	Paginate int64
	PerPage  int64
	Search   string
	Filters  []Filter
	Sort     Sort
	Date     time.Time
}

type Filter struct {
	Option    string `json:"option"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
	ValueType string `json:"type"`
}

type Sort struct {
	By   string `json:"by"`
	Type string `json:"type"`
}

// PageRequest : For Set Page Request
func PageRequest(param url.Values) (pageRequest *PageRequestStruct) {
	pageRequest = &PageRequestStruct{
		Filters:  make([]Filter, 0),
		Sort:     Sort{},
		Page:     1,
		PerPage:  10,
		Paginate: 0,
		Search:   "",
	}

	if param.Get("page") != "" {
		pageRequest.Page, _ = strconv.ParseInt(param.Get("page"), 0, 64)
	}

	if param.Get("per_page") != "" {
		pageRequest.PerPage, _ = strconv.ParseInt(param.Get("per_page"), 0, 64)
	}

	if param.Get("paginate") != "" {
		pageRequest.Paginate = 1
	}

	if param.Get("search") != "" {
		pageRequest.Search = param.Get("search")
	}

	if param.Get("date") != "" {
		layout := "2006-01-02"

		date, _ := time.Parse(layout, param.Get("date"))

		pageRequest.Date = date
	}

	if param.Get("sort") != "" {
		srts := strings.Split(param.Get("sort"), "|")
		if len(srts) >= 2 {
			pageRequest.Sort.By = srts[0]
			pageRequest.Sort.Type = srts[1]
		}
	}

	//Fetching filters data
	if len(param["filter[]"]) > 0 {
		for _, v := range param["filter[]"] {
			filter := Filter{}
			json.Unmarshal([]byte(v), &filter)
			pageRequest.Filters = append(pageRequest.Filters, filter)
		}
	}

	return

}
