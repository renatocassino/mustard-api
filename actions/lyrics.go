package actions

import (
	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/models"
)

type LyricsListDTO struct {
	Data []models.Lyric `json:"data"`
}

// GET /lyrics
func LyricListHandler(c echo.Context) error {
	user := GetUser(c)

	lyricsDTO := LyricsListDTO{}
	lyricsDTO.Data = models.Lyric{}.GetLyrics(user)

	return c.JSON(200, lyricsDTO)
}

func LyricCreateHandler(c echo.Context) error {
	user := GetUser(c)

	l := models.Lyric{}
	if err := c.Bind(&l); err != nil {
		return err
	}

	l.Create(user)
	return c.JSON(200, l)
}

func LyricUpdateHandler(c echo.Context) error {
	user := GetUser(c)

	id := c.Param("id")
	l := models.Lyric{}
	if err := c.Bind(&l); err != nil {
		return err
	}

	l.Update(id, user)
	return c.JSON(200, l)
}

func LyricDeleteHandler(c echo.Context) error {
	user := GetUser(c)

	id := c.Param("id")
	l := models.Lyric{}
	l.Delete(id, user)
	return c.JSON(204, "")
}
