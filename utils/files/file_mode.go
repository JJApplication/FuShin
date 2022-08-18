/*
Create: 2022/8/19
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package files
package files

// 常用的文件权限mode
// Usr uid
// G gid
const (
	Mask             = 0022
	ReadOnlyMode     = 0444
	WriteOnlyMode    = 0222
	UsrReadOnlyMode  = 0400
	UsrWriteOnlyMode = 0200
	GReadOnlyMode    = 0440
	GWriteOnlyMode   = 0220
	RWMode           = 0666
	UsrRWMode        = 0600
	GRWMode          = 0660
	XMode            = 0777
	UsrXMode         = 0700
	GXMode           = 0770

	FileMode  = 0640
	XFileMode = 0700
	DirMode   = 0750
	TmpMode   = 0660
)
