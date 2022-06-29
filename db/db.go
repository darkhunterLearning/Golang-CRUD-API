package db

import (
	"database/sql"
	"fmt"

	"github.com/darkhunterLearning/Golang-CRUD-API/action"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "Customer"
)

// DB set up
func Set_UpDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err := sql.Open("postgres", dbinfo)

	action.CheckErr(err)

	return DB
}
