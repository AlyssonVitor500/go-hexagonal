package db

import (
	"database/sql"
	"github.com/alyssonvitor500/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare(`select id, name, price, status from products where id=?`)

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	insert := `insert into products (id, name, price, status) values (?,?,?,?)`
	stmt, err := p.db.Prepare(insert)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	update := `update products set name = ?, price = ?, status = ? where id = ?`
	stmt, err := p.db.Prepare(update)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select count(id) from products where id=?", product.GetID()).Scan(&rows)

	if rows == 0 {
		return p.create(product)
	}

	return p.update(product)
}
