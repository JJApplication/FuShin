/*
Create: 2022/8/24
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

import (
	"fmt"
	"strconv"
)

// 文件大小计算

// CalcSize 计算一个文件大小得可读性值
func CalcSize(s int64) string {
	if s == 0 {
		return "0b"
	} else if s < KB {
		return fmt.Sprintf("%sb", strconv.FormatInt(s, 10))
	} else if s >= KB && s < MB {
		return fmt.Sprintf("%skb", strconv.FormatInt(s/KB, 10))
	} else if s >= MB && s < GB {
		return fmt.Sprintf("%smb", strconv.FormatInt(s/MB, 10))
	} else if s >= GB && s < TB {
		return fmt.Sprintf("%sgb", strconv.FormatInt(s/GB, 10))
	} else {
		return fmt.Sprintf("%stb", strconv.FormatInt(s/TB, 10))
	}
}
