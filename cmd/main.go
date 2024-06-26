package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/blazeisclone/order-service/domain/product"
	mysqlDB "github.com/blazeisclone/order-service/internal/database/mysql"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Env vars not loaded from file")
	}

	router := http.NewServeMux()
	product.Routes(router)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server listening on port:", port)

	db, err := mysqlDB.Init()
	if err != nil {
		fmt.Println("database init", err)
	}

	defer func() {
		db.Close()
		fmt.Println("db.Closed")
	}()

	migrate, err := mysqlDB.Migrate(db, &mysql.Config{})
	if err != nil {
		panic(error.Error(err))
	}

	row := db.QueryRowContext(context.Background(), "select * from products")

	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)

		if err := migrate.Up(); err != nil {
			panic(error.Error(err))
		}
	}

	server.ListenAndServe()
}
