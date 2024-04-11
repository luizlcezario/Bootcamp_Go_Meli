package handler

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
)

type ReqCreateUser struct {
	Name     string  `json:"name"`
	SourName string  `json:"sour_name"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Heigth   float32 `json:"height"`
}

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

func (se *HandlerUser) CreateNewUser(ctx *gin.Context) {
	jsonbBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "body is missing",
		})
	}
	var userReq ReqCreateUser
	if err := json.Unmarshal(jsonbBody, &userReq); err != nil {
		ctx.JSON(400, gin.H{
			"message": "not possible to parse the body verify json",
		})
	}
	user, err := se.service.Store(userReq.Name, userReq.SourName, userReq.Email, userReq.Age, userReq.Heigth)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(201, user)
}
