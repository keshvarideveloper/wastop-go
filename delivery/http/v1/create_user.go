package v1

import (
	"net/http"

	"github.com/keshvarideveloper/wastop/adapter/store"
	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/interactor/auth"
	"github.com/labstack/echo/v4"
)

func SignupUser(store store.MySQLStore, validator contract.ValidateSignupUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.SignupUserRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		resp, err := auth.New(store).SignupUser(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
