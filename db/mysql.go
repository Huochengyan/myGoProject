package db

import (
	"database/sql"
	"fmt"
)

const (
	DB_Driver = "root:root@tcp(localhost:3306)/hcydb?charset=utf8"
)

func openDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	CheckErr(err)
	return isOpen, db
}

var mysqlDB *sql.DB

func InitMysqlDb() *sql.DB {
	if mysqlDB == nil {
		result, DB := openDB()
		if result {
			mysqlDB = DB
		}
	}
	return mysqlDB
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
}
