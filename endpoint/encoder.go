package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	Data       interface{} `json:"data"`
	Errors     []string    `json:"errors"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Total uint `json:"total,omitempty"`
	Pages uint `json:"pages,omitempty"`
}

func PaginationData(totalEntries uint, totalPages uint) Pagination {
	return Pagination{
		Total: totalEntries,
		Pages: totalPages,
	}
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
