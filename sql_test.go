package godbmysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptsql := "INSERT INTO customer(id, name) VALUES('eko', 'Eko')"

	// Exec Context : untuk operasi SQL yang tidak membutuhkan hasil
	// Query Context :  function untuk melakukan query ke database menggunakan QueryContext(context, sql, params)
	_, err := db.ExecContext(ctx, scriptsql)

	if err != nil {
		panic(err)
	}

	println("Success INSERT new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptsql := "SELECT id, name FROM customers"

	rows, err := db.QueryContext(ctx, scriptsql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptsql := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customers"

	rows, err := db.QueryContext(ctx, scriptsql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, email string
		var name sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &createdAt, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		if name.Valid {
			fmt.Println("name :", name.String)
		}
		fmt.Println("email :", email)
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("birth_date :", birthDate)
		fmt.Println("created_at :", createdAt)
		fmt.Println("married :", married)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) { // BAHAYA SQL INJECTION
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin';#"
	password := "admin"

	scriptsql := "SELECT username FROM user WHERE username = '" + username + "'  AND password = '" + password + "' LIMIT 1"
	fmt.Println(scriptsql) // cek query
	rows, err := db.QueryContext(ctx, scriptsql)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		rows.Scan(username)
		fmt.Println("Sukses Login" + username)
	} else {
		fmt.Println("Gagal LOgin")
	}

	defer rows.Close()
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "richo"
	password := "richo123"

	scriptsql := "SELECT username FROM user WHERE username = ?  AND password = ? LIMIT 1"
	fmt.Println(scriptsql)                                           // cek query
	rows, err := db.QueryContext(ctx, scriptsql, username, password) // menambahkan parameter ketiga di QueryContext

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		rows.Scan(username)
		fmt.Println("Sukses Login" + username)
	} else {
		fmt.Println("Gagal LOgin")
	}

	defer rows.Close()
}

// versi Exec
func TestSqlExecSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "richo"
	password := "richo123"

	ctx := context.Background()

	scriptsql := "INSERT INTO user(username, password) VALUES(?, ?)"

	// Exec Context : untuk operasi SQL yang tidak membutuhkan hasil
	// Query Context :  function untuk melakukan query ke database menggunakan QueryContext(context, sql, params)
	_, err := db.ExecContext(ctx, scriptsql, username, password)

	if err != nil {
		panic(err)
	}

	println("Success INSERT new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "richo@gmail.com"
	comment := "halo saya berkomentar"

	scriptsql := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// Exec Context : untuk operasi SQL yang tidak membutuhkan hasil
	// Query Context :  function untuk melakukan query ke database menggunakan QueryContext(context, sql, params)
	result, err := db.ExecContext(ctx, scriptsql, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId() // untuk cek last id
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses menambahkan komentar", insertId)

	println("Success INSERT new customer")
}
