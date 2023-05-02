package app

import (
	"Dzaakk/go-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:kozato321@tcp(localhost:3306)/go_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate -database "mysql://root:kozato321@tcp(localhost:3306)/go_database_migration" -path db/migrations up  
	// migrate -database "mysql://root:kozato321@tcp(localhost:3306)/go_database_migration" -path db/migrations down 
}

