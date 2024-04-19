package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthReqired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := os.Getenv("USERAUTH")
		password := os.Getenv("PASSAUTH")
		usr, pass, ok := c.Request.BasicAuth()
		fmt.Println(user, usr, password, pass)
		if ok && usr == user && pass == password {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "please inform the correct credentials",
			})
		}

	}
}
