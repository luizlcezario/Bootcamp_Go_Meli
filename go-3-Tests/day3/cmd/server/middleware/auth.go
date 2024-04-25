package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luizlcezario/MelicampGoWeb/pkg/webpack"
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
			webpack.NewResponse(http.StatusMethodNotAllowed, nil, "Bad Creentials").Response(c)
		}

	}
}
