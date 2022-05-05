package v1

import (
	"net/http"

	"github.com/keshvarideveloper/wastop/adapter/store"
	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/interactor/user"
	"github.com/labstack/echo/v4"
)

func SignUpUser(store store.MySQLStore, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.CreateUserRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		resp, err := user.New(store).SignUpUser(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
