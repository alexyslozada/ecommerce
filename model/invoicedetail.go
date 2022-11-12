package model

import (
	"github.com/google/uuid"
)

// InvoiceDetail model of table invoice_details
type InvoiceDetail struct {
	ID        uuid.UUID `json:"id"`
	InvoiceID uuid.UUID `json:"invoice_id"`
	ProductID uuid.UUID `json:"product_id"`
	Amount    uint      `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

func (i InvoiceDetail) HasID() bool {
	return i.ID != uuid.Nil
}

// InvoiceDetails slice of InvoiceDetail
type InvoiceDetails []InvoiceDetail

func (i InvoiceDetails) IsEmpty() bool { return len(i) == 0 }
