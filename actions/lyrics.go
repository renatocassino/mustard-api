package actions

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/models"
)

type LyricsListDTO struct {
	Data []models.Lyric `json:"data"`
}

// OPTIONS /lyrics
func OptionsHandler(c echo.Context) error {
	header := c.Response().Header()
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Content-Length", "0")
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Max-Age", "1728000")
	return c.String(200, "")
}

// GET /lyrics
func LyricListHandler(c echo.Context) error {
	user := GetUser(c)

	fmt.Println("LYRICS GET")

	lyricsDTO := LyricsListDTO{}
	lyricsDTO.Data = models.Lyric{}.GetLyrics(user)

	if lyricsDTO.Data == nil {
		lyricsDTO.Data = []models.Lyric{}
	}

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
