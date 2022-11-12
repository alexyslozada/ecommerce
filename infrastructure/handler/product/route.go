package product

import (
	"github.com/alexyslozada/ecommerce/infrastructure/handler/middle"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/product"
	productStorage "github.com/alexyslozada/ecommerce/infrastructure/postgres/product"
)

// NewRouter returns a router to handle model.Product requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := product.New(productStorage.New(dbPool))
	return newHandler(useCase)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/admin/products", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/products")

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}
