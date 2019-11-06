package actions

import (
	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/models"
)

type LyricsListDTO struct {
	Data models.Lyrics `json:"data"`
}

// GET /lyrics
func LyricListHandler(c echo.Context) error {
	user := GetUser(c)

	lyricsDTO := LyricsListDTO{}
	lyricsDTO.Data = models.Lyric{}.GetLyrics(user)

	return c.JSON(200, lyricsDTO)
}
