package handler

import (
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"net/http"
)

const tmplPath = "src/template/"

// htmlBlob HTMLデータを返却する
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

// render
func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.HTMLBlob(http.StatusOK, b)
}
