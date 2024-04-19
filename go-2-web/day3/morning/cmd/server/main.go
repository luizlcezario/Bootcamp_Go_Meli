package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luizlcezario/MelicampGoWeb/cmd/server/handler"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
)

func main() {
	r := gin.Default()
	repo := users.NewRepository("user.json")
	service := users.NewService(repo)
	usersHandler := handler.NewHandlerUser(service)

	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Olá Luiz",
		})
	})
	gropuser := r.Group("/user")
	gropuser.GET("/all", usersHandler.ListAllFilter)
	gropuser.GET("/:id", usersHandler.FindByIdController)
	gropuser.POST("/create", usersHandler.CreateNewUser)
	gropuser.PUT("/change/:id", usersHandler.UpdateAllUser)
	gropuser.DELETE("/:id", usersHandler.DeleteUser)
	gropuser.PATCH("/change/partial/:id", usersHandler.ChangeNameAgeUser)
	if err := r.Run(); err == nil {
		fmt.Println("Server Online -->-->")
	}
}
