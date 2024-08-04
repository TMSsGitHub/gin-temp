package handler

import (
	"gin-temp/internal/global/errs"
	"gin-temp/internal/global/resp"
	"gin-temp/internal/service"
	"github.com/gin-gonic/gin"
)

// GetUserByPhone
// @Tags [user]
// @Summary 获取用户信息
// @Description 根据手机号获取用户信息
// @Accept json
// @Produce json
// @param token header string true "Token"
// @param phone path string true "手机号"
// @success 200 {object} resp.R
// @success 500 {object} resp.R
// @Router /user/{phone} [get]
func GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")
	if phone == "" {
		c.Set(resp.RES, resp.Fail("参数错误"))
		c.Abort()
		return
	}
	userService := service.GetUserService()
	user, err := userService.GetUserByPhone(phone)
	if err != nil {
		c.Error(errs.NewServerErr("获取用户信息失败", err))
		c.Abort()
		return
	}
	c.Set(resp.RES, resp.Success(user))
}
