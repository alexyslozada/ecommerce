package user

import (
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/user"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/response"
	"github.com/alexyslozada/ecommerce/model"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return h.responser.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.responser.Created(m))
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "useCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(users))
}
