package router

import (
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/robert/middelwares/handler"
	"github.com/robert/middelwares/middleware"
)

func Startserver() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", middleware.Authmiddleware(), handler.Mainpage)
	r.GET("/login", handler.Loginget)
	r.POST("/login", handler.Loginpost)
	r.Run()
}
