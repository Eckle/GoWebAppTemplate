package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pbentes/80_20/src/templates/components"
	"github.com/pbentes/80_20/src/templates/pages"
)

func Set(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.Layout(pages.Index))
	})

	router.GET("/content", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", pages.Content())
	})

	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", pages.Index())
	})
}
