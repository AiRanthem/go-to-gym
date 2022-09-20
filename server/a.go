package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var app = gin.Default()

func Run() {
	Api()
	err := app.Run()
	log.WithError(err).Warn("web server stopped")
}
