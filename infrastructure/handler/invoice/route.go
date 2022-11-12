package invoice

import (
	"github.com/alexyslozada/ecommerce/infrastructure/postgres/invoicereport"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/invoice"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/middle"
	invoiceStorage "github.com/alexyslozada/ecommerce/infrastructure/postgres/invoice"
)

// NewRouter returns a router to handle model.Invoice requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	// build middlewares to validate permissions on the routes
	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := invoice.New(invoiceStorage.New(dbPool), invoicereport.New(dbPool))

	return newHandler(useCase)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/admin/invoices", middlewares...)

	route.GET("", h.GetAll)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/private/invoices", middlewares...)

	route.GET("", h.MyShops)
}
