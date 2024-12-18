package main

import (
	"database/sql"
	sqlite3DBAdapter "github.com/alyssonvitor500/go-hexagonal/adapters/db"
	app "github.com/alyssonvitor500/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	defer db.Close()
	productDbAdapter := sqlite3DBAdapter.NewProductDb(db)
	productService := app.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example 1", 45)

	productService.Enable(product)
}
