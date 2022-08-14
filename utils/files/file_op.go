/*
Create: 2022/8/14
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/goutil/fsutil"
)

// 文件操作

// MkDir 按照子目录创建目录
func MkDir(p string, perm os.FileMode) error {
	return os.MkdirAll(p, perm)
}

// RmDir 删除目录
func RmDir(p string) error {
	return os.RemoveAll(p)
}

func RmFile(f string) error {
	return os.Remove(f)
}

// CreateFile 创建文件如果不存在
func CreateFile(f string, perm os.FileMode, dirPerm os.FileMode) (*os.File, error) {
	return fsutil.CreateFile(f, perm, dirPerm)
}

// Copy 复制文件到某路径
func Copy(src, dst string) error {
	return fsutil.CopyFile(src, dst)
}

// ClearFile 清空文件内容
func ClearFile(f string) error {
	return ioutil.WriteFile(f, []byte(""), 0644)
}

// Name 获取文件的base名称
func Name(f string) string {
	return fsutil.Name(f)
}

// TmpFile 创建/tmp下的临时文件
func TmpFile(f string) (*os.File, error) {
	tmp := filepath.Join(os.TempDir(), f)
	return os.Create(tmp)
}

// SetFileMode chmod修改文件权限
func SetFileMode(f string, perm os.FileMode) error {
	return os.Chmod(f, perm)
}

// SetFileOwn chown修改属主 user id group id
func SetFileOwn(f string, uid int, gid int) error {
	return os.Chown(f, uid, gid)
}

// GetContent 读取文件内容
// 错误时返回空
func GetContent(f string) []byte {
	if !IsExist(f) || !IsFile(f) {
		return nil
	}
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil
	}
	return data
}

// UnZip 解压文件
func UnZip(f string, target string) error {
	return fsutil.Unzip(f, target)
}

func IsZip(f string) bool {
	return fsutil.IsZipFile(f)
}

// Zip 压缩文件或目录到指定路径
// 例如: Zip(/tmp/1.txt, /tmp/1.zip)
// target不存在zip后缀时默认添加
func Zip(f string, target string) error {
	if !IsExist(f) {
		return errors.New("file or path not exist")
	}
	if !strings.HasSuffix(target, ".zip") {
		target += ".zip"
	}
	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}
	w := zip.NewWriter(targetFile)
	if IsFile(f) {
		src, _ := os.Open(f)
		defer src.Close()
		h := &zip.FileHeader{
			Name:   filepath.Base(f),
			Method: zip.Deflate,
		}
		fileName, _ := w.CreateHeader(h)
		io.Copy(fileName, src)
		_ = w.Flush()
	} else {
		filepath.WalkDir(f, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			src, _ := os.Open(path)
			defer src.Close()
			h := &zip.FileHeader{
				Name:   path,
				Method: zip.Deflate,
			}
			fileName, _ := w.CreateHeader(h)
			io.Copy(fileName, src)
			_ = w.Flush()
			return nil
		})
	}

	return w.Close()
}
