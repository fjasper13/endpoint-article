package database

import (
	"strconv"

	"github.com/fjasper13/endpoint-article/app/article/request"
)

func BuildQuery(pr *request.PageRequestStruct) (res string) {
	if pr.Search != "" || len(pr.Filters) > 0 {
		res += " WHERE "
	}
	if pr.Search != "" {
		res += "title LIKE '%" + pr.Search + "%' OR body LIKE '%" + pr.Search + "%'"
	}
	if len(pr.Filters) > 0 {
		if pr.Search != "" {
			res += " AND "
		}
		for i, v := range pr.Filters {
			if i == len(pr.Filters)-1 {
				res += v.Option + " " + v.Operator + " '" + v.Value + "' "
			} else {
				res += v.Option + " " + v.Operator + " '" + v.Value + "' AND "
			}
		}
	}
	if pr.Sort.By != "" && pr.Sort.Type != "" {
		res += " ORDER BY " + pr.Sort.By + " " + pr.Sort.Type
	} else {
		res += " ORDER BY created_at DESC "
	}
	if pr.Paginate == 1 {
		res += " LIMIT " + strconv.FormatInt(pr.PerPage, 10)
		res += " OFFSET " + strconv.FormatInt((pr.Page-1)*pr.PerPage, 10)
	}
	return
}
