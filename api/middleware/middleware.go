package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 注册跨域、统一响应体中间件
	r.Use(CorsMiddle(), Resp)
}
