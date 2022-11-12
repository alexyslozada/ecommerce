package response

import (
	"github.com/alexyslozada/ecommerce/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct{}

// New returns a new response API
func New() API {
	return API{}
}

func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡listo!"}},
	}
}

func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordCreated, Message: "¡listo!"}},
	}
}

func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}

func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordDeleted, Message: "¡listo!"}},
	}
}

func (a API) BindFailed(err error) error {
	e := model.NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}

func (a API) Error(c echo.Context, who string, err error) *model.Error {
	e := model.NewError()
	e.Err = err
	e.APIMessage = "¡Uy! metimos la pata, disculpanos lo solucionaremos"
	e.Code = UnexpectedError
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)
	// Only to avoid the panic error
	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
