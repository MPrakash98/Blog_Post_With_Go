package interfaces

import (
	_ "blog-post-api/models"
	"net/http"
)


type BlogPost interface {
	GetPosts(w http.ResponseWriter, r *http.Request) ()
	GetPostByID(w http.ResponseWriter, r *http.Request) ()
	CreatePost(w http.ResponseWriter, r *http.Request) ()
	DeletePost(w http.ResponseWriter, r *http.Request) ()
}