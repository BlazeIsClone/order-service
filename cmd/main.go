package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/blazeisclone/order-service/domain/product"
	"github.com/go-sql-driver/mysql"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order service reached\n"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handle)
	product.Routes(router)

	os.Setenv("PORT", "3000")

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server listening on port:", port)

	c := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		DBName:               "db",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("db.Closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	row := db.QueryRowContext(context.Background(), "select * from products")

	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	server.ListenAndServe()
}
