package routes

import (
	"blog-post-api/db"
	"blog-post-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func BlogRoute(router *mux.Router) {

	// Creates the db instance and passes it to api methods
	dbInstance := db.SetupDB()
	newInstance := services.CreateBlogPostStore(dbInstance)

	router.Use(commonMiddleware);

	// Get all posts
	router.HandleFunc("/posts/", newInstance.GetPosts).Methods("GET")
	
	// Get single post by id
	router.HandleFunc("/posts/{postID}", newInstance.GetPostByID).Methods("GET")

	// Create a post
	router.HandleFunc("/posts/", newInstance.CreatePost).Methods("POST")

	// Delete a specific post by the postID
	router.HandleFunc("/posts/{postID}", newInstance.DeletePost).Methods("DELETE")

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

