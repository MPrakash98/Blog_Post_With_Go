package db

import (
	"blog-post-api/helper"
	"blog-post-api/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", models.DB_USER, models.DB_PASSWORD, models.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	fmt.Println(err)

	helper.CheckErr(err)

	return db
}
