package function

import (
	"os"
	"path/filepath"
)

// 获取绝对路径
func GetAbsPath(relativePath string) string {
	execPath, _ := os.Executable()
	path, _ := filepath.Split(execPath)
	if relativePath == "" {
		return ""
	}
	return filepath.Join(path, relativePath)
}

// 获取当前执行程序绝对路径（不兼容于go run main.go运行模式）
func GetCurrentAbsPath() string {
	execPath, _ := os.Executable()
	path, _ := filepath.Split(execPath)
	return path
}

//检测路径不存在则创建

func FileNotExistsAndCreate(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}
