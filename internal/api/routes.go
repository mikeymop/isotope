package api

import (
	"github.com/isotope/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddRoutes(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: web.BuildHTTPFS(),
		// Root:       "web",
		HTML5: true,
	}))

	api := e.Group("/api")

	api.GET("/photos", handleListAlbums())
	api.GET("/photos/:album", handleListPhotos())
	api.GET("/photos/:album/:photo", handleGetPhoto())
}
