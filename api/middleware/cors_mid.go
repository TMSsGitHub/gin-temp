package middleware

import (
	"fmt"
	"gin-temp/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 生产环境不会填* 应当指定域名 fixme
			c.Header("Access-Control-Allow-Origin", "*")
			// 允许使用的HTTP METHOD
			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE, OPTIONS")
			// 允许使用的请求头
			c.Header("Access-Control-Allow-Headers", fmt.Sprintf("Origin, X-Requested-With, Content-Type, Accept, Authorization, %s", conf.Cfg.App.AuthKey))
			// 允许客户端的响应头
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			// 是否需要携带认证信息，可以是 cookies; authorization headers; TLS client certificates
			// 设为true时，Access-Control_Allow-Origin不能为 * ??
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 放行OPTION请求，但不执行后续方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
