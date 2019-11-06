package actions

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/models"
)

type LyricsListDTO struct {
	data models.Lyrics `json:"data"`
}

// GET /lyrics
func LyricListHandler(c echo.Context) error {
	user := GetUser(c)
	return c.String(http.StatusOK, fmt.Sprintf("FOI BONITO FOI %s", user.GivenName))
}
