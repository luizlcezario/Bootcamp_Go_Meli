package handler

import (
	"github.com/gin-gonic/gin"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
)

type HandlerUser struct {
	service users.Service
}

func NewHandlerUser(se users.Service) *HandlerUser {
	return &HandlerUser{
		service: se,
	}
}

// ex01 afternoom

func (se *HandlerUser) ListAllFilter(ctx *gin.Context) {
	result, err := se.service.GetAllAndFilter(ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(404, gin.H{"message": err.Error()})
	}
	ctx.JSON(200, result)
}

// ex02  afternoom

func (se *HandlerUser) FindByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if result, err := se.service.FindById(id); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, result)
	}
}
