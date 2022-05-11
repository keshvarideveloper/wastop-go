package v1

import (
	"net/http"
	"strconv"

	"github.com/keshvarideveloper/wastop/adapter/store"
	"github.com/keshvarideveloper/wastop/contract"
	"github.com/keshvarideveloper/wastop/dto"
	"github.com/keshvarideveloper/wastop/interactor/user"
	"github.com/labstack/echo/v4"
)

func UpdateProfile(store store.MySQLStore, validator contract.ValidateUpdateProfile) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		userID, err := strconv.Atoi(idStr)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var req = dto.UpdateProfileRequest{}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		req.ID = uint(userID)
		if err := validator(c.Request().Context(), req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		resp, err := user.New(store).UpdateProfile(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)

	}
}
