package godbmysql

import (
	"context"
	"fmt"
	"testing"
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

	scriptsql := "SELECT id, name FROM customer"

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
