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

// ListUsers godoc
// @Summary List users
// @Tags User
// @Description get users
// @securityDefinitions.basic BasicAuth
// @Accept json
// @Produce json
// @Success 202 {object} webpack.Response
// @Failure 404 {object} webpack.Response
// @Router /user/all [get]
func (se *HandlerUser) ListAllFilter(ctx *gin.Context) {
	result, err := se.service.GetAllAndFilter(ctx.Request.URL.Query())
	if err != nil {
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusAccepted, result, "").Response(ctx)
}

// FindUserId
// @Summary find User
// @Tags User
// @Description get User
// @Accept json
// @Produce json
// @Param id path string true "id do user"
// @Success 202 {object} webpack.Response
// @Failure 404 {object} webpack.Response
// @Router /user/:id [get]
func (se *HandlerUser) FindByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if result, err := se.service.FindById(id); err != nil {
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	} else {
		webpack.NewResponse(http.StatusAccepted, result, "").Response(ctx)
	}
}

// PutUser godoc
// @Summary update user
// @Tags User
// @Description find user by id and update the fields
// @Param request body handler.ReqCreateUser true "body template"
// @Accept json
// @Produce json
// @Success 201 {object} webpack.Response
// @Failure 400 {object} webpack.Response
// @Router /user/change/:id [put]
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

// CreateUser godoc
// @Summary create user
// @Tags User
// @Description Create user using a entity
// @securityDefinitions.basic BasicAuth
// @Param request body handler.ReqCreateUser true "body template"
// @Accept json
// @Produce json
// @Success 201 {object} webpack.Response
// @Failure 400 {object} webpack.Response
// @Router /user/create [post]
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

// DeleteUser godoc
// @Summary delete user
// @Tags User
// @Description find a user by id and try to delete it from the store
// @Param id path string true "id do user"
// @Accept json
// @Produce json
// @Success 204 {object} webpack.Response
// @Failure 404 {object} webpack.Response
// @Router /user/:id [delete]
func (se *HandlerUser) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := se.service.Delete(id); err != nil {
		webpack.NewResponse(http.StatusNotFound, nil, err.Error()).Response(ctx)
	}
	webpack.NewResponse(http.StatusNoContent, nil, "").Response(ctx)
}

// ChangePartial godoc
// @Summary patch user
// @Tags User
// @Description find a user by id and try to delete it from the store
// @Param request body handler.ReqCreateUser true "id do user"
// @Param id path string true "id do user"
// @Accept json
// @Produce json
// @Success 201 {object} webpack.Response
// @Failure 400 {object} webpack.Response
// @Router /user/:id [patch]
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
