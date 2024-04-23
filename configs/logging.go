package configs

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() {

	gin.DisableConsoleColor()

	path, _ := os.Getwd()
	newPath := filepath.Join(path + "/storages/logs")
	os.MkdirAll(newPath, os.ModePerm)

	f, _ := os.Create(newPath + "/gin-" + time.Now().Format("20060102150405") + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}
