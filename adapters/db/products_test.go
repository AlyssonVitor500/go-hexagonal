package db_test

import (
	"database/sql"
	"github.com/alyssonvitor500/go-hexagonal/adapters/db"
	"github.com/alyssonvitor500/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `create table products ( 
    			"id" string, 
				"name" string, 
				"price" float, 
				"status" string
                      );`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createProduct(db *sql.DB) {
	insert := `insert into products (id, name, price, status) values ("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDB := db.NewProductDb(Db)
	product, err := productDB.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()
	productDB := db.NewProductDb(Db)

	product := application.NewProduct()

	// Saving product
	product.Name = "New Product Test"
	product.Price = 42.0

	err := product.Enable()
	require.Nil(t, err)

	productResult, err := productDB.Save(product)
	require.Nil(t, err)

	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	// Updating product
	product.Name = "Updated Product Test"
	product.Price = 0.0

	err = product.Disable()
	require.Nil(t, err)

	productResult, err = productDB.Save(product)
	require.Nil(t, err)

	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}
