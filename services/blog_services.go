package services

import (
	"blog-post-api/helper"
	"blog-post-api/interfaces"
	"blog-post-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogPostStore struct {
	BlogPostDB *sql.DB
}

func CreateBlogPostStore(db *sql.DB) interfaces.BlogPost {
	return &BlogPostStore{
		BlogPostDB: db,
	}
}

// Get all posts

func (BP *BlogPostStore) GetPosts(w http.ResponseWriter, r *http.Request) {
	helper.PrintMessage("Getting posts...")

	// Get all posts from posts table
	rows, err := BP.BlogPostDB.Query("SELECT * FROM posts")

	// check errors
	helper.CheckErr(err)

	// var response []models.JsonResponse
	var posts []models.Post

	// Foreach post
	for rows.Next() {
		var id int
		var postID string
		var post string

		err = rows.Scan(&id, &postID, &post)

		// check errors
		helper.CheckErr(err)

		posts = append(posts, models.Post{PostID: postID, Post: post})
	}

	var response = models.JsonResponse{Type: "success", Data: posts}
	json.NewEncoder(w).Encode(response)
	//return response, nil
}

// Get a post by ID

func (BP *BlogPostStore) GetPostByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	postID := params["postID"]

	helper.PrintMessage("Getting posts...")

	// Get all posts from posts table having postID = postId
	rows, err := BP.BlogPostDB.Query("SELECT * FROM posts where postID = $1", postID)

	// check errors
	helper.CheckErr(err)

	// var response []models.JsonResponse
	var posts models.Post

	// Foreach post
	for rows.Next() {
		var id int
		var postID string
		var post string

		err = rows.Scan(&id, &postID, &post)

		// check errors
		helper.CheckErr(err)

		posts = models.Post{PostID: postID, Post: post}
	}

	var response = models.JsonShortResponse{Type: "success", Data: posts}

	json.NewEncoder(w).Encode(response)

}

// Create a post

func (BP *BlogPostStore) CreatePost(w http.ResponseWriter, r *http.Request) {
	// postID := r.Context().Value("postID")
	// post := r.PostFormValue("post")

	postID := r.PostFormValue("postID")
	post := r.PostFormValue("post")
	// fmt.Println(post)
	var response = models.JsonResponse{}

	if postID == "" || post == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing postID or post parameter."}
	} else {

		helper.PrintMessage("Inserting post into DB")

		fmt.Println("Inserting new post with ID: " + postID + " and name: " + post)

		var lastInsertID int
		err := BP.BlogPostDB.QueryRow("INSERT INTO posts(postID, post) VALUES($1, $2) returning id;", postID, post).Scan(&lastInsertID)

		// check errors
		helper.CheckErr(err)

		response = models.JsonResponse{Type: "success", Message: "The post has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a post

func (BP *BlogPostStore) DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID := params["postID"]

	var response = models.JsonResponse{}

	if postID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing postID parameter."}
	} else {

		helper.PrintMessage("Deleting post from DB")

		_, err := BP.BlogPostDB.Exec("DELETE FROM posts where postID = $1", postID)

		// check errors
		helper.CheckErr(err)

		response = models.JsonResponse{Type: "success", Message: "The post has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
