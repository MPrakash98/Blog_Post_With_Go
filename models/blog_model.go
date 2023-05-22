package models

type Post struct {
	PostID string `json:"postID"`
	Post   string `json:"post"`
}
type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Post `json:"data"`
	Message string `json:"message"`
}

type JsonShortResponse struct {
	Type    string `json:"type"`
	Data    Post   `json:"data"`
	Message string `json:"message"`
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "m@password"
	DB_NAME     = "postgres"
)
