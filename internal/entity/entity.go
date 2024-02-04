package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID: uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CategoryId string `json:"category_id"`
	ImageURL string `json:"image_url"`
}

func NewProduct( name, description string, price float64, categoryId, imageURL string) *Product{
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Description: description,
		Price: price,
		CategoryId: categoryId,
		ImageURL: imageURL,
	}
}

func (p *Product) ValidateProduct() error {
	if p.Name == "" {
		return errors.New("product name should not be empty")
	}

	if p.Price == 0 {
		return errors.New("product price should not be empty")
	}

	if p.CategoryId == "" {
		return errors.New("product categoryId should not be empty")
	}

	return nil
}

func (c *Category) ValidaCategory() error {
	if c.Name == "" {
		return errors.New("product name should not be empty")
	}

	return nil
}