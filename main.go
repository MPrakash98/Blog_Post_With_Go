package main

import (
	_ "blog-post-api/db"
	"blog-post-api/routes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Init the mux router
	router := mux.NewRouter()
	//db.SetupDB()
	routes.BlogRoute(router) //add this
	// Route handles & endpoints

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]string{"data": "Hello from Blog Post"})
	}).Methods("GET")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
