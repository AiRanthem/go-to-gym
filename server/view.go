package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func View() {
	app.LoadHTMLFiles("./static/index.html")
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
