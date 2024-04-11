package handler

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func GetBody[T any](body T, ctx *gin.Context) (T, error) {
	jsonbBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return body, errors.New("body is missing")
	}
	if err := json.Unmarshal(jsonbBody, &body); err != nil {
		return body, errors.New("not possible to parse the body verify json")
	}
	return body, validator.Validate(body)

}
