package middleware

import (
	"gin-temp/conf"
	"gin-temp/internal/global/errs"
	"gin-temp/internal/global/resp"
	"gin-temp/internal/global/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get(conf.Cfg.App.AuthKey)
	if token == "" {
		c.Error(errs.SimpleErrWithCode(resp.NeedToLogin, "请先登录"))
		c.Abort()
		return
	}

	claims, err := utils.ValidateAccessToken(token)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "token is expired") {
			c.JSON(resp.LoginExpired, resp.Fail("access已过期"))
			c.Abort()
			return
		}
		// 其他解析错误
		c.Error(errs.SimpleErrWithCode(resp.NeedToLogin, "请重新登录"))
		c.Abort()
		return
	}
	c.Set("claims", claims) // todo 用途为定
	c.Next()
}
