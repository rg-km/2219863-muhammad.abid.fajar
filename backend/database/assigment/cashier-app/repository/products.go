package repository

import (
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id

	var product Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products WHERE id = ?"

	row, err := p.db.Query(sqlStatement, id)

	if err != nil {
		return product, err
	}

	defer row.Close()

	for row.Next() {
		if err := row.Scan(&product.ID, &product.Category, &product.ProductName, &product.Price); err != nil {
			return product, err
		}
	}

	return product, nil

	// return Product{}, nil // TODO: replace this
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name

	var product Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products WHERE product_name = ?"

	row := p.db.QueryRow(sqlStatement, productName)

	err := row.Scan(&product.ID, &product.Category, &product.ProductName, &product.Price, &product.Quantity)

	if err != nil {
		return product, err
	}

	return product, nil
	// return Product{}, nil // TODO: replace this
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products

	var products []Product
	var sqlStatement string

	sqlStatement = "SELECT * FROM products"

	rows, err := p.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
	// return []Product{}, nil // TODO: replace this
}
