package paypal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/invoice"
	"github.com/alexyslozada/ecommerce/domain/paypal"
	"github.com/alexyslozada/ecommerce/domain/purchaseorder"

	storageInvoice "github.com/alexyslozada/ecommerce/infrastructure/postgres/invoice"
	storagePurchaseOrder "github.com/alexyslozada/ecommerce/infrastructure/postgres/purchaseorder"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	purchaseOrderUseCase := purchaseorder.New(storagePurchaseOrder.New(dbPool))
	invoiceUseCase := invoice.New(storageInvoice.New(dbPool), nil)
	useCase := paypal.New(purchaseOrderUseCase, invoiceUseCase)

	return newHandler(useCase)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/paypal")

	route.POST("", h.Webhook)
}
