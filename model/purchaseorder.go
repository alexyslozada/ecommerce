package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// PurchaseOrder model of table purchase_orders
type PurchaseOrder struct {
	ID        uuid.UUID       `json:"id"`
	UserID    uuid.UUID       `json:"user_id"`
	Products  json.RawMessage `json:"products"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
}

func (p PurchaseOrder) HasID() bool {
	return p.ID != uuid.Nil
}

func (p PurchaseOrder) Validate() error {
	if len(p.Products) == 0 {
		return errors.New("the list of products can't be empty")
	}

	var ptps []ProductToPurchase
	err := json.Unmarshal(p.Products, &ptps)
	if err != nil {
		return fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	for _, v := range ptps {
		if v.ProductID == uuid.Nil {
			return errors.New("the product id can´t be empty")
		}
		if v.Amount < 1 {
			return errors.New("the amount of products can't be less than 1")
		}
		if v.UnitPrice < 0 {
			return errors.New("the unit price can't be negative")
		}
	}

	return nil
}

func (p PurchaseOrder) TotalAmount() float64 {
	if len(p.Products) == 0 {
		return 0
	}

	var ptps []ProductToPurchase
	err := json.Unmarshal(p.Products, &ptps)
	if err != nil {
		return 0
	}

	var total float64
	for _, v := range ptps {
		total += float64(v.Amount) * v.UnitPrice
	}

	return total
}

// PurchaseOrders slice of PurchaseOrder
type PurchaseOrders []PurchaseOrder

func (p PurchaseOrders) IsEmpty() bool { return len(p) == 0 }

type ProductToPurchase struct {
	ProductID uuid.UUID `json:"product_id"`
	Amount    uint      `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
}

type ProductToPurchases []ProductToPurchase
