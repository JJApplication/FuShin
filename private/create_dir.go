/*
Create: 2022/8/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package private

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/JJApplication/fushin/pkg"
	"github.com/JJApplication/fushin/utils/json"
)

// 创建服务的目录

// CreateTag 生成服务的标签
func createTag(projectName, Type, DataType string, protocol []string) {
	fmt.Println("start to create project meta")
	data, err := json.Json.MarshalIndent(map[string]interface{}{
		"create":    time.Now(),
		"generator": pkg.Fushin,
		"project":   projectName,
		"type":      Type,
		"metaType":  DataType,
		"protocol":  protocol,
	}, "", "  ")
	if err != nil {
		fmt.Printf("create tag error: %s\n", err.Error())
		return
	}
	// 保证在当前目录下的projectName/projectName.meta.json
	err = ioutil.WriteFile(path.Join(projectName, fmt.Sprintf("%s.meta.json", projectName)), data, 0644)
	if err != nil {
		fmt.Printf("create tag error: %s\n", err.Error())
		return
	}
}

func CreateDir(projectName, Type, DataType string, protocol []string) {
	fmt.Println("start to create project")
	if projectName == "" {
		return
	}
	err := os.Mkdir(projectName, 0644)
	if err != nil {
		fmt.Printf("create project dir error: %s\n", err.Error())
		return
	}
	createTag(projectName, Type, DataType, protocol)
}
