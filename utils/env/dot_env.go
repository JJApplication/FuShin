/*
Create: 2023/2/19
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package env

import (
	"os"

	"github.com/joho/godotenv"
)

// 支持dotenv
// 可以与EnvLoader一起使用

const (
	DEFAULT    = ".env"
	DEV        = "development"
	PROD       = "production"
	DEVFILE    = ".env.dev"
	PRODFILE   = ".env.prod"
	GO_DEVMODE = "GO_DEVMODE"
)

// Load 默认读取当前目录下的.env
func Load(files ...string) error {
	return godotenv.Load(files...)
}

// AutoLoad 默认读取当前目录下的.env
// 在不同开发模式时自动匹配
// 内置了dev prod模式分别会匹配.env.dev和.env.prod 根据当前环境变量GO_DEVMODE判断
func AutoLoad() {
	mode := os.Getenv(GO_DEVMODE)
	switch mode {
	case DEV:
		_ = godotenv.Load(DEVFILE)
	case PROD:
		_ = godotenv.Load(PRODFILE)
	default:
		_ = godotenv.Load(DEFAULT)
	}
}

// Read 尝试从环境变量中读取合并的环境变量到Map
// var myEnv map[string]string
// myEnv, err := Read()
//
// var1 := myEnv["Var1"]
func Read() (map[string]string, error) {
	return godotenv.Read()
}

// Unmarshal 尝试解析字符串到环境变量Map
// Unmarshal("KEY=VALUE")
func Unmarshal(s string) (map[string]string, error) {
	return godotenv.Unmarshal(s)
}

func UnmarshalBytes(b []byte) (map[string]string, error) {
	return godotenv.UnmarshalBytes(b)
}

func Marshal(envs map[string]string) (string, error) {
	return godotenv.Marshal(envs)
}

// Write 存储Map到文件中
func Write(envs map[string]string, file string) error {
	return godotenv.Write(envs, file)
}

// Overload 覆写已经存在的环境变量
func Overload(files ...string) error {
	return godotenv.Overload(files...)
}
