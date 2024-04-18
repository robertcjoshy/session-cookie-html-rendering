package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mainpage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Loginget(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Loginpost(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "r" && password == "1" {
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"message": "invalid email or password"})
	}
}
