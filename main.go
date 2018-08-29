package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gittokkunn/go-jwt/jwt"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./public/index.html", "./public/mypage.html", "./public/signup.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/signup", jwt.SignUp)
	r.GET("/adduser", jwt.AddUser)
	r.GET("/mypage", jwt.LoginMyPage)
	r.Run(":3003")
}
