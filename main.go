package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tacnoman/mustard-api/actions"
	"github.com/tacnoman/mustard-api/core"
	"github.com/tacnoman/mustard-api/models"
)

func isAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		authorization := request.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			return echo.ErrUnauthorized
		}

		tokenstring := strings.Replace(authorization, "Bearer ", "", 1)
		user, err := models.User{}.LoggedUser(tokenstring)
		if err != nil {
			return echo.ErrUnauthorized
		}

		c.Set("user", user)

		return next(c)
	}
}

func main() {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(core.GetEnv("JWT_TOKEN", "secret")),
	})

	e := echo.New()
	e.GET("/api/v1/rhymes/:language/:word", actions.RhymesHandler)
	e.GET("/auth", actions.AuthHandler)
	e.GET("/auth/callback", actions.AuthCallbackHandler)
	e.GET("/api/v1/lyrics", actions.LyricListHandler, authMiddleware)

	e.Logger.Fatal(e.Start(":8000"))

	// m := martini.Classic()
	// m.Use(render.Renderer())

	// m.Get("/api/v1/lyrics", actions.LyricList)
	// m.Get("/api/v1/rhymes/:language/:word", actions.RhymesHandler)
	// m.Run()
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
