package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Product model of table products
type Product struct {
	ID          uuid.UUID       `json:"id"`
	ProductName string          `json:"product_name"`
	Price       float64         `json:"price"`
	Images      json.RawMessage `json:"images"`
	Description string          `json:"description"`
	Features    json.RawMessage `json:"features"`
	CreatedAt   int64           `json:"created_at"`
	UpdatedAt   int64           `json:"updated_at"`
}

func (p Product) HasID() bool {
	return p.ID != uuid.Nil
}

// Products slice of Product
type Products []Product

func (p Products) IsEmpty() bool { return len(p) == 0 }
