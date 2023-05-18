package main

import (
	"database/sql"

	db2 "github.com/JoaoLeonello/hex-arch-studies/adapters/db"
	"github.com/JoaoLeonello/hex-arch-studies/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemple", 30)

	productService.Enable(product)
}
