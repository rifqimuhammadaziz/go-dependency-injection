package app

import (
	"database/sql"
	"time"

	"github.com/rifqimuhammadaziz/go-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3360)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
