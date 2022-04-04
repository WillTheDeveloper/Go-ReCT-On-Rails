package main

import (
	"database/sql"
	"os"
)

func connectDatabase() {
	db := sql.Open("mysql", os.Getenv("db_user")+":"+os.Getenv("db_pass")+"@tcp("+os.Getenv("db_url")+":3306/"+os.Getenv("db_name")+"")
}
