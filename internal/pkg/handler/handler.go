package handler

import (
	"api/pkg/infrastucture/db"
)

type HTTPHandler struct {
}

func NewHTTPHandler(db db.Database) *HTTPHandler {

	return &HTTPHandler{}
}
