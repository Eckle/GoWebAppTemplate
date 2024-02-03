package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pbentes/Chime/src/templates/components"
	"github.com/pbentes/Chime/src/templates/pages"
	"github.com/pbentes/Chime/src/utils"
)

func Set(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, components.Layout(pages.Index))
	})
}
