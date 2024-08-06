package service

import (
	"crypto/sha256"
	"fmt"
	"gin-temp/conf"
	"gin-temp/internal/global/errs"
	"gin-temp/internal/global/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
)

type FileService struct{}

var (
	fileService *FileService
	fileOnce    sync.Once
)

func GetFileService() *FileService {
	fileOnce.Do(func() {
		fileService = &FileService{}
	})
	return fileService
}

func (*FileService) FileUpload(file *multipart.FileHeader) (string, error) {
	// 创建临时文件
	tempFile, err := file.Open()
	if err != nil {
		return "", errs.NewServerErr("文件错误", err)
	}
	defer tempFile.Close()
	//计算SHA-256
	hashed := sha256.New()
	if _, err := io.Copy(hashed, tempFile); err != nil {
		return "", errs.NewServerErr("文件错误", err)
	}
	fileHash := hashed.Sum(nil)
	fileHashStr := fmt.Sprintf("%x", fileHash)
	prefix := fileHashStr[:8]
	nowTs := utils.GetCurrentTs()
	filename := file.Filename
	fileType := filepath.Ext(filename)
	newFilename := fmt.Sprintf("%s%d%s", prefix, nowTs, fileType)
	// 查表存不存在相同的哈希
	// 存在则更新引用次数
	// 不存在则新增记录并保存到本地

	// 重新打开文件句柄，确保在复制到目标路径时使用新鲜的句柄
	if _, err := tempFile.Seek(0, io.SeekStart); err != nil {
		return "", errs.NewServerErr("打开文件时发生错误", err)
	}
	// 构建完整的保存路径
	targetPath := filepath.Join(conf.Cfg.File.Dir, newFilename)
	// 创建目标目录如果不存在
	if err := os.MkdirAll(filepath.Dir(targetPath), 755); err != nil {
		return "", errs.NewServerErr("文件上传失败", err)
	}
	// 创建文件
	dst, err := os.Create(targetPath)
	if err != nil {
		return "", errs.NewServerErr("上传文件失败", err)
	}
	defer dst.Close()
	// 将上传的文件内容写入新创建的文件
	if _, err := io.Copy(dst, tempFile); err != nil {
		return "", errs.NewServerErr("上传文件时发生错误", err)
	}
	return newFilename, nil
}

func (*FileService) FileDownload(filePath string) ([]byte, string, error) {
	// 打开文件
	file, err := os.Open(fmt.Sprintf("%s/%s", conf.Cfg.File.Dir, filePath))
	if err != nil {
		return nil, "", errs.NewServerErr("找不到文件了", err)
	}
	defer file.Close()
	// 获取文件名
	filename := filepath.Base(filePath)
	// 读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, "", errs.NewServerErr("读取文件失败", err)
	}
	return fileContent, filename, nil
}
