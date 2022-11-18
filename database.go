package godbmysql

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:dev034@@tcp(localhost:3306)/belajar_go_sql?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)  // jika dalam 5 menit tidak ada transaksi maka akan di close
	db.SetConnMaxLifetime(60 * time.Minute) // set 60 menit koneksi apapun akan diperbarui dengan koneksi baru

	return db
}
