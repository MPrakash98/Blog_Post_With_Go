package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"blog-post-api/db"
	"blog-post-api/helper"
	"blog-post-api/models"

	"github.com/gorilla/mux"
)

// Get all posts

// response and request handlers
func GetPosts(w http.ResponseWriter, r *http.Request) {
	db := db.SetupDB()

	helper.PrintMessage("Getting posts...")

	// Get all posts from posts table that don't have postID = "1"
	rows, err := db.Query("SELECT * FROM posts")

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
}

// Get a post by ID

// response and request handlers
func GetPostByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	postID := params["postID"]

	db := db.SetupDB()
	helper.PrintMessage("Getting posts...")

	// Get all posts from posts table that don't have postID = "1"
	rows, err := db.Query("SELECT * FROM posts where postID = $1", postID)

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

// response and request handlers
func CreatePost(w http.ResponseWriter, r *http.Request) {
	postID := r.PostFormValue("postID")
	post := r.PostFormValue("post")
	// fmt.Println(post)
	var response = models.JsonResponse{}

	if postID == "" || post == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing postID or post parameter."}
	} else {
		db := db.SetupDB()

		helper.PrintMessage("Inserting post into DB")

		fmt.Println("Inserting new post with ID: " + postID + " and name: " + post)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO posts(postID, post) VALUES($1, $2) returning id;", postID, post).Scan(&lastInsertID)

		// check errors
		helper.CheckErr(err)

		response = models.JsonResponse{Type: "success", Message: "The post has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a post

// response and request handlers
func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID := params["postID"]

	var response = models.JsonResponse{}

	if postID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing postID parameter."}
	} else {
		db := db.SetupDB()

		helper.PrintMessage("Deleting post from DB")

		_, err := db.Exec("DELETE FROM posts where postID = $1", postID)

		// check errors
		helper.CheckErr(err)

		response = models.JsonResponse{Type: "success", Message: "The post has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
