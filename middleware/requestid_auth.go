package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/requestid"
	"kang/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// RequestIdAuth requestId中间件
func RequestIdAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		//开始时间
		reqStartTime := time.Now().UnixMilli()

		//处理请求
		c.Next()

		if c.Request.Method != "GET" {
			c.Request.ParseForm()
		}

		var responseData string
		//结束时间
		reqEndTime := time.Now().UnixMilli()

		if c.Writer.Status() == 200 {
			responseData = writer.body.String()
			global.G_Logger.Warn(
				responseData,
				zap.String("req_body", c.Request.PostForm.Encode()),
				zap.String("req_host", c.Request.Host),
				zap.String("req_method", c.Request.Method),
				zap.String("req_clientIp", c.ClientIP()),
				zap.String("req_id", requestid.Get(c)),
				zap.String("req_uri", c.Request.RequestURI),
				zap.String("req_time", fmt.Sprintf("%vms", reqEndTime-reqStartTime)),
			)
		} else {
			global.G_Logger.Info(
				responseData,
				zap.String("req_body", c.Request.PostForm.Encode()),
				zap.String("req_host", c.Request.Host),
				zap.String("req_method", c.Request.Method),
				zap.String("req_clientIp", c.ClientIP()),
				zap.String("req_id", requestid.Get(c)),
				zap.String("req_uri", c.Request.RequestURI),
				zap.String("req_time", fmt.Sprintf("%vms", reqEndTime-reqStartTime)),
			)
		}
	}
}
