package paypal

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/alexyslozada/ecommerce/model"
)

type UseCase interface {
	ProcessRequest(header http.Header, body []byte) error
}

type UseCasePurchaseOrder interface {
	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}

type UseCaseInvoice interface {
	Create(m *model.PurchaseOrder) error
}
