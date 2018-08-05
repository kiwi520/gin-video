package main

import (
	"gin-video/router"
	_"gin-video/database"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func init()  {

	if gin.Mode() == "ReleaseMode" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{})
	}
}

//192.168.33.10
func main() {

	r := router.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
