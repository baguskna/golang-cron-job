package app

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func NewDB() {
	dbMsql, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_cron_job?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	dbMsql.SetConnMaxIdleTime(5)
	dbMsql.SetMaxOpenConns(20)
	dbMsql.SetConnMaxLifetime(60 * time.Minute)
	dbMsql.SetConnMaxIdleTime(10 * time.Minute)

	db = dbMsql

	fmt.Println("DB connected")
}

func GetDB() *sql.DB {
	return db
}
