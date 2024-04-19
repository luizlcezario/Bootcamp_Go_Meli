package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
	"github.com/luizlcezario/MelicampGoWeb/pkg/webpack"
)

type ReqCreateUser struct {
	Name     string  `json:"name" validate:"min=2,max=30,regexp=^[a-zA-Z]*$"`
	SourName string  `json:"sour_name" validate:"min=2,max=30,regexp=^[a-zA-Z]*$"`
	Email    string  `json:"email" validate:"min=6"`
	Age      int     `json:"age" validate:"min=18"`
	Heigth   float32 `json:"height" validate:"min=1.0"`
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
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusAccepted, result, "").Response(ctx)
}

// ex02  afternoom

func (se *HandlerUser) FindByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if result, err := se.service.FindById(id); err != nil {
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	} else {
		webpack.NewResponse(http.StatusAccepted, result, "").Response(ctx)
	}
}

func (se *HandlerUser) UpdateAllUser(ctx *gin.Context) {
	var userReq ReqCreateUser
	id := ctx.Param("id")
	userReq, err := GetBody(userReq, ctx)
	if err != nil {
		webpack.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	user, err := se.service.UpdateUser(id, userReq.Name, userReq.SourName, userReq.Email, userReq.Age, userReq.Heigth)
	if err != nil {
		webpack.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusCreated, user, "").Response(ctx)
}

func (se *HandlerUser) CreateNewUser(ctx *gin.Context) {
	var userReq ReqCreateUser
	userReq, err := GetBody(userReq, ctx)
	if err != nil {
		webpack.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	user, err := se.service.Store(userReq.Name, userReq.SourName, userReq.Email, userReq.Age, userReq.Heigth)
	if err != nil {
		webpack.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusCreated, user, "").Response(ctx)
}

func (se *HandlerUser) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := se.service.Delete(id); err != nil {
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusNoContent, nil, "").Response(ctx)
}

func (se *HandlerUser) ChangeNameAgeUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var body struct {
		SourName string `json:"sour_name" validate:"min=2,max=30,regexp=^[a-zA-Z]*$"`
		Age      int    `json:"age" validate:"min=18"`
	}
	GetBody(body, ctx)
	user, err := se.service.UpdateUser(id, "", body.SourName, "", body.Age, 0.0)
	if err != nil {
		webpack.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusCreated, user, "").Response(ctx)
}
