package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/actions"
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
		_, err := models.User{}.LoggedUser(tokenstring)
		if err != nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func main() {
	// authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(core.GetEnv("JWT_TOKEN", "secret")),
	// })

	e := echo.New()
	e.GET("/api/v1/rhymes/:language/:word", actions.RhymesHandler)
	e.GET("/auth", actions.AuthHandler)
	e.GET("/auth/callback", actions.AuthCallbackHandler)
	e.GET("/api/v1/lyrics", actions.LyricListHandler, isAuthorized)
	e.POST("/api/v1/lyrics", actions.LyricCreateHandler, isAuthorized)
	e.PUT("/api/v1/lyrics/:id", actions.LyricUpdateHandler, isAuthorized)
	e.DELETE("/api/v1/lyrics/:id", actions.LyricDeleteHandler, isAuthorized)
	e.Logger.Fatal(e.Start(":8000"))
}
