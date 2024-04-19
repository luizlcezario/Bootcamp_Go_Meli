package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
)

func generatedId(s string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x\n", bt)
	}
}
func main() {
	r := gin.Default()
	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Olá Luiz",
		})
	})
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Users)
	})
	r.Run()
}
