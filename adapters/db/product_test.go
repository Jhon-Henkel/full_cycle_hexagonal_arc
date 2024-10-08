package db_test

import (
	"database/sql"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/adapters/db"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (id string, name string, price float, status string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products (id, name, price, status) values ("abc", "Product Test", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDB.Save(product)
	if err != nil {
		log.Fatal(err.Error())
	}

	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDB.Save(product)
	if err != nil {
		log.Fatal(err.Error())
	}

	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}
