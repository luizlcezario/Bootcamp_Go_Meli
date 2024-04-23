package webpack

import "github.com/gin-gonic/gin"

type Response struct {
	Code uint        `json:"status"`
	Data interface{} `json:"data,omitempty"`
	Erro string      `json:"error_message,omitempty"`
}

func NewResponse(code uint, data interface{}, erro string) Response {
	return Response{
		code, data, erro,
	}
}

func (e Response) Response(c *gin.Context) {
	c.AbortWithStatusJSON(int(e.Code), e)
}
