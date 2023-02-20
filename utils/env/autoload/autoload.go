/*
Create: 2023/2/21
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package autoload 旁载后自动加载.env文件
package autoload

import "github.com/joho/godotenv"

func init() {
	godotenv.Load()
}
