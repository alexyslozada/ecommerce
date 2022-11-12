package product

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/alexyslozada/ecommerce/domain/product"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/response"
	"github.com/alexyslozada/ecommerce/model"
)

type handler struct {
	useCase  product.UseCase
	response response.API
}

func newHandler(useCase product.UseCase) handler {
	return handler{useCase: useCase}
}

// Create handles the creation of a model.Product
func (h handler) Create(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

// Update handles the update of a model.Product
func (h handler) Update(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}
	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

// Delete handles the deleting of a model.Product
func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(productData))
}

// GetAll handles the search of all model.Product
func (h handler) GetAll(c echo.Context) error {
	products, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(products))
}
