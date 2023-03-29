/*
Create: 2023/3/28
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"errors"
	"os"
)

// 判断缓存证书目录下是否存在证书

var certCacheDir string

// SetCertCacheDir 设置全局的证书缓存目录
func SetCertCacheDir(p string) {
	certCacheDir = p
}

// 在为设置缓存目录时报错
func isCertDir() bool {
	return certCacheDir == ""
}

func HasCertCache() bool {
	if !isCertDir() {
		return false
	}
	if _, err := os.Stat(certCacheDir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CleanCertCache() error {
	if !isCertDir() {
		return errors.New(errNoCert)
	}
	if _, err := os.Stat(certCacheDir); os.IsNotExist(err) {
		return nil
	}
	return os.RemoveAll(certCacheDir)
}

func CreateCertDir() error {
	if !isCertDir() {
		return errors.New(errNoCert)
	}
	if _, err := os.Stat(certCacheDir); os.IsNotExist(err) {
		return os.MkdirAll(certCacheDir, 0644)
	}
	return nil
}
