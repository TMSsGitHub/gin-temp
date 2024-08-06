package handler

import (
	"fmt"
	"gin-temp/conf"
	"gin-temp/internal/global/constant"
	"gin-temp/internal/global/errs"
	"gin-temp/internal/global/resp"
	"gin-temp/internal/service"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

// FileUpload
// @Tags [file]
// @Summary 文件上传
// @Description 将文件上传至服务器
// @Accept multipart/form-data
// @Produce json
// @param file formData file true "需要上传的文件"
// @success 200 {object} resp.R "文件上传成功"
// @success 500 {object} resp.R "文件上传失败"
// @Router /file/upload [post]
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(errs.NewServerErr("出错了", err))
		c.Abort()
		return
	}

	fileService := service.GetFileService()
	filename, err := fileService.FileUpload(file)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	url := fmt.Sprintf("http://%s:%d/file/download/%s", conf.Cfg.App.Host, conf.Cfg.App.Port, filename)
	c.Set(resp.RES, url)
}

// FileDownload
// @Tags [file]
// @Summary 文件下载
// @Description 根据文件名下载文件
// @Accept json
// @Produce json
// @param url path string true "文件名"
// @success 200 {object} resp.R
// @success 500 {object} resp.R
// @Router /file/download/{url} [get]
func FileDownload(c *gin.Context) {
	url := c.Param("url")
	fileService := service.GetFileService()
	file, fileName, err := fileService.FileDownload(url)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	fileType := strings.ToLower(filepath.Ext(fileName))
	c.Data(200, constant.MimeTypes[fileType], file)
}
