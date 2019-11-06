package actions

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/models"
)

func GetUser(c echo.Context) models.User {
	authorization := c.Request().Header.Get("Authorization")
	tokenstring := strings.Replace(authorization, "Bearer ", "", 1)
	user, _ := models.User{}.LoggedUser(tokenstring)

	return *user
}
