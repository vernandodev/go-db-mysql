package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	go_db_mysql "github.com/vernandodev/go-db-mysql"
	"github.com/vernandodev/go-db-mysql/models"
)

func TestCommentInsert(t *testing.T) {
	commentrepository := NewCommentRepository(go_db_mysql.GetConnection())

	ctx := context.Background()

	comment := models.Comment{
		Email:   "test@gmail.com",
		Comment: "Test comment ke 2",
	}

	result, err := commentrepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_mysql.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_mysql.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
