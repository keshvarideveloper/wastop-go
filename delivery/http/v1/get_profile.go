package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/keshvarideveloper/wastop/adapter/store"
	"github.com/keshvarideveloper/wastop/dto"
	user "github.com/keshvarideveloper/wastop/interactor/users"
	"github.com/labstack/echo/v4"
)

func GetProfile(store store.MySQLStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		fmt.Println(idStr)
		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var req = dto.GetProfileRequest{ID: uint(userID)}

		// if err := validator(c.Request().Context(), req); err != nil {
		// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		// }

		resp, err := user.New(store).GetProfile(c.Request().Context(), req)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
