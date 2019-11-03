package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/tacnoman/mustard-api/actions"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/api/v1/lyrics", actions.LyricList)
	m.Get("/api/v1/rhymes/:language/:word", actions.RhymesHandler)
	m.Run()
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
