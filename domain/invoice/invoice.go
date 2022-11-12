package invoice

import (
	"github.com/alexyslozada/ecommerce/model"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error
}

type Storage interface {
	Create(m *model.Invoice, ms model.InvoiceDetails) error
}
