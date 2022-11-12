package handler

import (
	"github.com/alexyslozada/ecommerce/infrastructure/handler/login"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/infrastructure/handler/product"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/purchaseorder"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/user"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)

	// A
	// B
	// C

	// L
	login.NewRouter(e, dbPool)

	// P
	product.NewRouter(e, dbPool)
	purchaseorder.NewRouter(e, dbPool)

	// U
	user.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"service_name": "",
			},
		)
	})
}
