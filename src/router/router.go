package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pbentes/80_20/src/templates/components"
	"github.com/pbentes/80_20/src/templates/pages"
	"github.com/pbentes/80_20/src/utils"
)

func Set(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, components.Layout(pages.Index))
	})
}
