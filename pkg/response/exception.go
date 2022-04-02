package response

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Status    int         `json:"status"`
	ErrCode   Code        `json:"errcode"`
	RequestId string      `json:"requestid"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// ResponseJson 基础返回
func ResponseJson(c *gin.Context, status int, errcode Code, msg string, data interface{})  {
	if msg == "" {
		msg =CodeMap[errcode]
	}
	c.JSON(status,JsonResponse{
		Status:    status,
		ErrCode:   errcode,
		RequestId: requestid.Get(c),
		Message:   msg,
		Data:      data,
	})
}

// NotFoundException 404错误
func NotFoundException(c *gin.Context, msg string) {
	if msg == "" {
		msg = CodeMap[RequestMethodErr]
	}

	ResponseJson(c,http.StatusBadRequest,RequestParamErr,msg,nil)
}
