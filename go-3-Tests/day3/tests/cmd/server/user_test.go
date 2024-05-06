package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luizlcezario/MelicampGoWeb/cmd/server/handler"
	"github.com/luizlcezario/MelicampGoWeb/internal/domain"
	users "github.com/luizlcezario/MelicampGoWeb/internal/user"
	"github.com/luizlcezario/MelicampGoWeb/pkg/webpack"
	"github.com/luizlcezario/MelicampGoWeb/tests/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	UpdateUser = "/user/change/:id"
)

func createProductHandler(t *testing.T) (*gin.Engine, *handler.HandlerUser, *mocks.MockStore) {
	t.Helper()
	server := CreateServer()
	mock := new(mocks.MockStore)
	repository := users.NewRepository(mock)
	serviceMock := users.NewService(repository)
	handler := handler.NewHandlerUser(serviceMock)
	return server, handler, mock
}

func TestUpdateUser(t *testing.T) {
	mockedUser := []domain.User{
		{Id: "test"}, {Id: "test2"},
	}

	newMocked := handler.ReqCreateUser{
		Name:     "Luiz",
		SourName: "Lima",
		Email:    "luiz@gmail.com",
		Age:      20,
		Heigth:   7.8,
	}
	resultData := domain.User{
		Id:       "test",
		Name:     "name",
		SourName: "lima",
	}
	t.Run("should return ok", func(t *testing.T) {
		g, ser, mockR := createProductHandler(t)
		var result webpack.Response
		var data domain.User
		paramId := "test"

		g.PUT(UpdateUser, ser.UpdateAllUser)
		mockR.On("Read").Return(mockedUser, nil)
		mockR.On("Write", resultData).Return(nil)
		body, e := json.Marshal(newMocked)
		assert.Nil(t, e)
		request, response := MakeRequest(http.MethodPut, fmt.Sprintf("/user/change/%v", paramId), string(body))

		g.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		err := json.Unmarshal(response.Body.Bytes(), &result)
		assert.Nil(t, err)
		jsonData, err := json.Marshal(result.Data)
		assert.Nil(t, err)
		err = json.Unmarshal(jsonData, &data)
		assert.Nil(t, err)
		assert.Equal(t, resultData, data)
		assert.True(t, mockR.AssertExpectations(t))
	})

}
