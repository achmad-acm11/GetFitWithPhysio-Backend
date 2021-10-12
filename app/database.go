package app

import (
	"GetfitWithPhysio-backend/helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConfigDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/getfitwith_physio")

	// Handle If Error
	helper.HandleError(err)

	return db
}
