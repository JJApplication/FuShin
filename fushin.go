/*
Create: 2022/7/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/JJApplication/fushin/inner"
	"github.com/JJApplication/fushin/private"
)

func exit(err error) {
	if err == terminal.InterruptErr {
		fmt.Println("exit")
		os.Exit(0)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("👻 Thanks for using Fushin to create application of projectJJ")
	fmt.Printf("🐼 Copyright: %s\n🔗 Go to site: http://%s\n📧 Email: %s\n",
		inner.COPYRIGHT,
		inner.SITE,
		inner.MAIL,
	)
	// 创建项目名称
	projectName := ""
	promptProjectName := &survey.Input{
		Message: "your project name: ",
	}
	exit(survey.AskOne(promptProjectName, &projectName))

	// 选择项目类型 【web app】 【noengine app】 【empty app】
	chooseType := ""
	promptChooseType := &survey.Select{
		Message: "Choose application type:",
		Options: private.APPType,
		Default: private.APPTypeDefault,
	}
	exit(survey.AskOne(promptChooseType, &chooseType))

	// 选择.octopus数据类型 【json】 【yaml】 【pig】
	chooseDataType := ""
	promptChooseDataType := &survey.Select{
		Message: "Choose metadata type:",
		Options: private.APPMetaType,
		Default: private.APPMetaTypeDefault,
	}
	exit(survey.AskOne(promptChooseDataType, &chooseDataType))

	// 选择服务通信协议 【http】 【http-based rpc】 【uds】
	var protocol []string
	promptProto := &survey.MultiSelect{
		Message: "What protocols do you prefer:",
		Options: private.APPProto,
		Default: private.APPProtoDefault,
	}
	exit(survey.AskOne(promptProto, &protocol))

	// 开始生成
	startTodo := false
	promptStart := &survey.Confirm{
		Message: "Do you like to create the project right now?",
	}
	exit(survey.AskOne(promptStart, &startTodo))
	if startTodo {
		private.CreateDir(projectName, chooseType, chooseDataType, protocol)
	}
	fmt.Println("🥤 Enjoy yourself!!!")
}
