package common

import (
	"errors"
	"log"
	"os"
	"strings"
)

func isExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func isDir(path string) bool {
	if !isExist(path) {
		return false
	}
	stat, _ := os.Stat(path)
	return stat.IsDir()
}

func FilePathCheck(path string) {
	if !isExist(path) {
		log.Fatal(errors.New("文件路径不存在"))
	}
	if isDir(path) {
		log.Fatal(errors.New("输入路径为文件夹"))
	}
}

func ModeSet(path string) {
	FilePathCheck(path)
	if strings.HasSuffix(path, ".db") {
		Mode = 1
	} else {
		Mode = 0
	}
}
