package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var words = []string{}

type LinksDTO struct {
	Self     *string `json:"self"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Last     *string `json:"last"`
}

type DataDTO struct {
	Language string   `json:"language"`
	Words    []string `json:"words"`
}

type RhymesDTO struct {
	Links LinksDTO `json:"links"`
	Data  DataDTO  `json:"data"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RhymesHandler(params martini.Params, req *http.Request, r render.Render) {
	if len(words) == 0 {
		file, err := ioutil.ReadFile("pt-br.txt")
		if err != nil {
			log.Fatal(err)
		}

		words = strings.Split(string(file), "\n")
	}

	qs := req.URL.Query()
	page, err := strconv.Atoi(qs.Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(qs.Get("limit"))
	if err != nil {
		limit = 30
	}

	fmt.Println(limit)
	fmt.Println(page)

	word := params["word"]
	part := word[len(word)-3:]

	rhymes := []string{}
	for _, w := range words {
		if strings.HasSuffix(w, part) {
			rhymes = append(rhymes, w)
		}
	}

	offset := (page - 1) * limit

	language := params["language"]
	total := len(rhymes)

	dto := RhymesDTO{
		Links: LinksDTO{
			Self:     nil,
			Previous: nil,
			Next:     nil,
			Last:     nil,
		},
		Data: DataDTO{
			Language: language,
			Words:    rhymes[offset:min(offset+limit, total)],
		},
	}

	if page > 0 {
		previous := fmt.Sprintf("/api/v1/rhymes/%s/%s?page=%d&limit=%d", language, word, page-1, limit)
		dto.Links.Previous = &previous
	}

	if offset+limit < total {
		next := fmt.Sprintf("/api/v1/rhymes/%s/%s?page=%d&limit=%d", language, word, page+1, limit)
		dto.Links.Next = &next
	}

	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	last := fmt.Sprintf("/api/v1/rhymes/%s/%s?page=%d&limit=%d", language, word, lastPage, limit)
	dto.Links.Last = &last

	self := fmt.Sprintf("/api/v1/rhymes/%s/%s?page=%d&limit=%d", language, word, page, limit)
	dto.Links.Self = &self

	r.JSON(200, dto)
}
