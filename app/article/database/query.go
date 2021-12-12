package database

import (
	"strconv"

	"github.com/fjasper13/endpoint-article/app/article/request"
)

func BuildQuery(pr *request.PageRequestStruct) (res string) {
	res += " WHERE deleted_at IS NULL "
	if pr.Search != "" {
		res += "AND title LIKE '%" + pr.Search + "%' OR body LIKE '%" + pr.Search + "%'"
	}
	if len(pr.Filters) > 0 {
		for _, v := range pr.Filters {
			res += " AND " + v.Option + " " + v.Operator + " '" + v.Value + "' "
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
