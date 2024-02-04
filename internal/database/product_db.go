package database

import (
	"database/sql"

	"github.com/analopesdev/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB{
	return &ProductDB{
		db: db,
	}
}

func (pd *ProductDB) GetProducts() ([] *entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, categoryId, imageURL FROM products")

  if err != nil {
		return nil, err 
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var  product entity.Product
		 err = rows.Scan(&product.ID, &product.Name,	&product.Description, &product.Price, &product.CategoryId, &product.ImageURL)

		 if err != nil {
			 return nil, err
		 }

		 products = append(products, &product)
		}

		return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, categoryId, imageURL) VALUES (?, ?, ?, ?, ?, ?)", product.CategoryId, product.Name, product.Description, product.CategoryId, product.ImageURL)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}


func(pd *ProductDB ) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price,
		 &product.CategoryId, &product.ImageURL)

		 if err != nil {
			return nil, err
		 }

		 return &product, nil
}


func (pd *ProductDB) GetProductByCategoryId(categoryId string) ([] *entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, categoryId, imageURL FROM products WHERE categoryId = ?", categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.ImageURL)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)

	}

	return products, nil
}