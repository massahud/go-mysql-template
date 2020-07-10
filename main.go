package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/massahud/retry"
)

func main() {

	timeout, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	result := retry.Func(timeout, 500*time.Millisecond, func(_ context.Context) (interface{}, error) {
		db, err := sql.Open("mysql", "root:********@tcp(db:3306)/hello")
		if err != nil {
			return nil, err
		}
		_, err = db.Query("select 1 from persons")
		if err != nil {
			return nil, err
		}
		fmt.Println("DB READY")
		return db, nil
	})
	if result.Err != nil {
		panic(result.Err.Error())
	}
	db := result.Value.(*sql.DB)
	defer db.Close()
	rows, err := db.Query("select * from persons")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println("Hello,", name)
	}

}
