package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luizlcezario/MelicampGoWeb/cmd/server/handler"
	"github.com/luizlcezario/MelicampGoWeb/cmd/server/middleware"
	"github.com/luizlcezario/MelicampGoWeb/docs"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mli Bootcamp Api
// @version 1.0
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	repo := users.NewRepository("user.json")
	service := users.NewService(repo)
	usersHandler := handler.NewHandlerUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Olá Luiz",
		})
	})
	gropuser := r.Group("/user")
	gropuser.GET("/all", middleware.AuthReqired(), usersHandler.ListAllFilter)
	gropuser.GET("/:id", usersHandler.FindByIdController)
	gropuser.POST("/create", middleware.AuthReqired(), usersHandler.CreateNewUser)
	gropuser.PUT("/change/:id", usersHandler.UpdateAllUser)
	gropuser.DELETE("/:id", usersHandler.DeleteUser)
	gropuser.PATCH("/change/partial/:id", usersHandler.ChangeNameAgeUser)
	if err := r.Run(); err == nil {
		fmt.Println("Server Online -->-->")
	}
}
