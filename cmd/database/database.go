package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

func InitDB(credentials string) error {
	var err error
	DB, err = sql.Open("mysql", credentials) // DO NOT TRY TO BE SMART AND *CLEAN UP* THIS ASSIGNMENT BY USING :=
	if err != nil {
		return fmt.Errorf("could not connect to DB, %v", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return nil
}
