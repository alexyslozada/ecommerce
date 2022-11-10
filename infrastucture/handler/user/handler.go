package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/user"
	"github.com/alexyslozada/ecommerce/model"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.Create(&m); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, m)
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
