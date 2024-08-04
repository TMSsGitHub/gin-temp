package logger

import (
	"bufio"
	"fmt"
	"gin-temp/conf"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type loggerWriter struct {
	mu        sync.Mutex
	dir       string
	ticker    *time.Ticker
	stopFlush chan bool
	file      map[string]*os.File
	bufWrs    map[string]*bufio.Writer
}

var Logger *zap.Logger

func newLoggerWriter(dir string, flushInterval int) *loggerWriter {
	lw := &loggerWriter{
		dir:       dir,
		ticker:    time.NewTicker(time.Duration(flushInterval) * time.Second),
		stopFlush: make(chan bool),
		file:      make(map[string]*os.File),
		bufWrs:    make(map[string]*bufio.Writer),
	}
	go lw.flushLoop()
	return lw
}

func (lw *loggerWriter) flushLoop() {
	defer lw.ticker.Stop()

	for {
		select {
		case <-lw.stopFlush:
			return
		case <-lw.ticker.C:
			lw.mu.Lock()
			for _, bufWr := range lw.bufWrs {
				bufWr.Flush()
			}
			lw.mu.Unlock()
		}
	}
}

func (lw *loggerWriter) Close() error {
	lw.stopFlush <- true
	lw.mu.Lock()
	defer lw.mu.Unlock()

	var errors []error
	for _, bufWr := range lw.bufWrs {
		if err := bufWr.Flush(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, file := range lw.file {
		if err := file.Close(); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (lw *loggerWriter) Write(p []byte) (n int, err error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()

	now := time.Now().Format("200601/02")
	filename := fmt.Sprintf("%s/%s.log", lw.dir, now)
	file, fileExists := lw.file[filename]
	if !fileExists {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// 创建目录
			dir := filepath.Dir(filename)
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println("创建日至目录失败：", err)
				return 0, err
			}
		}
	}
	// 打开文件
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return 0, err
	}
	lw.file[filename] = file

	_, err = os.Stdout.Write(p)
	if err != nil {
		return 0, err
	}
	bufWr, bufWrExists := lw.bufWrs[filename]
	if !bufWrExists {
		bufWr = bufio.NewWriter(file)
		lw.bufWrs[filename] = bufWr
	}
	return bufWr.Write(p)
}

func (lw *loggerWriter) Sync() error {
	lw.mu.Lock()
	defer lw.mu.Unlock()

	for _, bufWr := range lw.bufWrs {
		if err := bufWr.Flush(); err != nil {
			return err
		}
	}
	for _, file := range lw.file {
		if err := file.Sync(); err != nil {
			return err
		}
	}
	return nil
}

func InitLogger() {
	writer := newLoggerWriter(conf.Cfg.Logger.Dir, 30)

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Sampling:         nil,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.AddSync(writer),
		config.Level,
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		//Logger.Info(fmt.Sprintf("| %15s | %6s | %s", clientIP, method, path))
		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		//clientUserAgent := c.Request.UserAgent()
		Logger.Info(fmt.Sprintf("| %5d | %9v | %15s | %6s | %s", statusCode, latency, clientIP, method, path))
	}
}
