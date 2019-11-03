package actions

import (
	"github.com/tacnoman/mustard-api/models"
)

type LyricsListDTO struct {
	data models.Lyrics `json:"data"`
}

// List gets all Lyrics. This function is mapped to the path
// // GET /lyrics
// func LyricList(params martini.Params, req *http.Request, r render.Render) error {
// 	lyrics := &models.Lyrics{}
// }
