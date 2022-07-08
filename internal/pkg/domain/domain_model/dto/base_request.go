package dto

type BaseRequest struct {
	Query     string                 `json:"query" form:"query"`
	Variables map[string]interface{} `json:"variables" form:"variables"`
}
