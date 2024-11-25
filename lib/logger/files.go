package logger

import (
	"fmt"
	"os"
)

// checkNotExist 函数检查给定路径是否不存在
func checkNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// checkPermission 检查给定路径是否有权限访问
func checkPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

// isNotExistMkDir 检查路径是否存在，如果不存在则创建目录
func isNotExistMkDir(src string) error {
	if notExist := checkNotExist(src); notExist == true {
		if err := mkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// mkDir 创建指定路径的目录
func mkDir(src string) error {
	// 创建目录，os.ModePerm 赋予所有权限
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// mustOpen 打开指定的文件，并确保目录存在
func mustOpen(fileName, dir string) (*os.File, error) {
	perm := checkPermission(dir)
	if perm == true {
		return nil, fmt.Errorf("permission denied dir: %s", dir)
	}

	err := isNotExistMkDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error during make dir %s, err: %s", dir, err)
	}
	// 尝试打开或创建文件，追加写入模式
	f, err := os.OpenFile(dir+string(os.PathSeparator)+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open file, err: %s", err)
	}

	return f, nil
}
