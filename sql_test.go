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

	// Exec : untuk operasi SQL yang tidak membutuhkan hasil
	// Sedangkan, function untuk melakukan query ke database menggunakan QueryContext(context, sql, params)
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
