package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tacnoman/mustard-api/actions"
	"github.com/tacnoman/mustard-api/core"
)

func main() {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(core.GetEnv("JWT_TOKEN", "secret")),
	})

	e := echo.New()
	e.GET("/api/v1/rhymes/:language/:word", actions.RhymesHandler)
	e.GET("/auth", actions.AuthHandler)
	e.GET("/auth/callback", actions.AuthCallbackHandler)
	e.GET("/api/v1/lyrics", actions.LyricListHandler, authMiddleware)
	e.POST("/api/v1/lyrics", actions.LyricCreateHandler, authMiddleware)
	e.PUT("/api/v1/lyrics/:id", actions.LyricUpdateHandler, authMiddleware)
	e.DELETE("/api/v1/lyrics/:id", actions.LyricDeleteHandler, authMiddleware)
	e.Logger.Fatal(e.Start(":8000"))
}
