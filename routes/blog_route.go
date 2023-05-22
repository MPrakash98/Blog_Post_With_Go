package routes

import (
	"blog-post-api/controllers" //add this

	"github.com/gorilla/mux"
)

func BlogRoute(router *mux.Router) {

	// Get all posts
	router.HandleFunc("/posts/", controllers.GetPosts).Methods("GET")

	// Get single post by id
	router.HandleFunc("/posts/{postID}", controllers.GetPostByID).Methods("GET")

	// Create a post
	router.HandleFunc("/posts/", controllers.CreatePost).Methods("POST")

	// Delete a specific post by the postID
	router.HandleFunc("/posts/{postID}", controllers.DeletePost).Methods("DELETE")

}
