package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"kang/global"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost: "localhost:443",
		})

		err := middleware.Process(c.Writer,c.Request)
		if err != nil {
			global.G_Logger.Error("HTTPS Request err: " + err.Error())
			return
		}
		//继续往下处理
		c.Next()
	}
}
