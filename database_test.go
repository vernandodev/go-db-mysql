package godbmysql

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:dev034@@tcp(localhost:3306)/db1")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // minimal
	db.SetMaxOpenConns(100)                 // maksimal
	db.SetConnMaxIdleTime(5 * time.Minute)  // dalam 5 menit tidak ada koneksi maka akan dimatikan
	db.SetConnMaxLifetime(60 * time.Minute) // koneksi apapun dalam 60 menit akan diperbarui

}
