package api

import (
	"github.com/isotope/internal/api/photos"
	"github.com/labstack/echo/v4"
	"net/http"
)

const libraryDir = "test"

func handleListAlbums() echo.HandlerFunc {
	return echo.HandlerFunc(
		func(c echo.Context) error {
			if c.Request().Method == "GET" {
				p := photos.ListAlbums(libraryDir)
				return c.JSON(http.StatusOK, p)
			}

			if c.Request().Method == "POST" {
				return c.String(http.StatusOK, "getAlbums expects GET")
			}
			return c.String(http.StatusOK, "Hello World")
		})
}

func handleGetAlbums() echo.HandlerFunc {
	return echo.HandlerFunc(
		func(c echo.Context) error {
			if c.Request().Method == "GET" {
				p := photos.ListAlbums("test")
				return c.JSON(http.StatusOK, p)
			}

			if c.Request().Method == "POST" {
				return c.String(http.StatusOK, "getAlbums expects GET")
			}
			return c.String(http.StatusOK, "Hello World")
		})
}

func handleListPhotos() echo.HandlerFunc {
	return echo.HandlerFunc(
		func(c echo.Context) error {
			album := c.Param("album")
			p := photos.ListPhotos(album)
			return c.JSON(http.StatusOK, p)
		})
}

func handleGetPhoto() echo.HandlerFunc {
	return echo.HandlerFunc(
		func(c echo.Context) error {
			album := c.Param("album")
			photo := c.Param("photo")
			response := photos.GetPhoto(album, photo)
			return c.JSON(http.StatusOK, response)
		})
}
