/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"fmt"
	"math"
	"os"
	"time"
)

// 文件状态

const (
	KB = 1 << 10
	MB = KB << 10
	GB = MB << 10
	TB = GB << 10
)

// IsExist 文件或文件夹是否存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile 是否为文件
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsFolder 是否为文件夹
func IsFolder(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetSize 获取文件大小
func GetSize(f string) int64 {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return 0
	}
	return info.Size()
}

// GetSizeReadable 可读性更好的fileSize
func GetSizeReadable(f string) string {
	size := GetSize(f)
	if size/TB > 0 {
		return fmt.Sprintf("%.2fTB", math.Floor(float64(size/TB)))
	} else if size/GB > 0 {
		return fmt.Sprintf("%.2fGB", math.Floor(float64(size/GB)))
	} else if size/MB > 0 {
		return fmt.Sprintf("%.2fMB", math.Floor(float64(size/MB)))
	} else if size/KB > 0 {
		return fmt.Sprintf("%.2fKB", math.Floor(float64(size/KB)))
	} else {
		return fmt.Sprintf("%dB", size)
	}
}

// GetModTime 获取文件修改时间
// 文件不存在时返回当前时间
func GetModTime(f string) time.Time {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return time.Now()
	}
	return info.ModTime()
}

// GetFileMode 获取文件类型
func GetFileMode(f string) os.FileMode {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return os.ModeType
	}
	return info.Mode()
}
