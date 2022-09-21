package server

import (
	"github.com/gin-gonic/gin"
	"go-to-gym/gym"
	"net/http"
)

func Api() {
	router := app.Group("/api")

	router.GET("/result", func(c *gin.Context) {
		tf := c.Query("TimeFunc")
		timeFunc, err := gym.GetTimeFunc(tf)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		ctx, cancel := gym.GetContext()
		defer cancel()
		pic, err := gym.GoToGym(ctx, timeFunc)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.File(pic)
	})
}
