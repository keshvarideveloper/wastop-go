package v1

import (
	"net/http"

	"github.com/keshvarideveloper/wastop/adapter/store"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/interactor/auth"
	"github.com/labstack/echo/v4"
)

func LoginUser(store store.MySQLStore) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.LoginUserWithEmailRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// if err := validator(req); err != nil {
		// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		// }

		resp, err := auth.New(store).LoginUserWithEmail(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
