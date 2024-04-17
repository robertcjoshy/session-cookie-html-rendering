package router

import (
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/robert/middelwares/handler"
)

func Startserver() {
	r := gin.Default()
	r.GET("/login", handler.Login)
}
