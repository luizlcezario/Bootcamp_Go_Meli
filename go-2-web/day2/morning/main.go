package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"slices"

	"github.com/gin-gonic/gin"
)

func generatedId(s string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x", bt)
	}
}

// ex01 afternoom

func GetValueByJsonField(field int, values []string, result []interface{}) []interface{} {
	return slices.DeleteFunc(result, func(e any) bool {
		value := reflect.ValueOf(e).Elem().Field(field - 1)
		var final string
		switch value.Type().Kind() {
		case reflect.Int:
			final = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			final = fmt.Sprintf("%v", value.Bool())
		case reflect.Float32:
			final = fmt.Sprintf("%0.2f", value.Float())
		case reflect.String:
			final = value.String()
		}
		for _, i := range values {
			if i == final {
				return false
			}
		}
		return true
	})
}

func listAllFilter(ctx *gin.Context) {
	querys := ctx.Request.URL.Query()
	result := make([]interface{}, len(Users))
	for i, user := range Users {
		result[i] = user
	}
	tags := Users[0].GetQueryParametersValids()
	for key, i := range querys {
		if field := tags[key]; field == 0 {
			ctx.JSON(404, gin.H{"error": "please send just query parameters valids"})
			return
		} else {
			result = GetValueByJsonField(field, i, result)
		}
	}
	ctx.JSON(200, result)
}

// ex02  afternoom

func findByIdController(ctx *gin.Context) {
	id := ctx.Param("id")
	if result, err := findById(Users, id); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, result)
	}
}
func main() {
	r := gin.Default()
	r.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Olá Luiz",
		})
	})
	gropuser := r.Group("/user")
	{
		gropuser.GET("/all", listAllFilter)
		gropuser.GET("/:id", findByIdController)
	}
	r.Run()
}
