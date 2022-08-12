/*
Create: 2022/8/12
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package string
package string

import (
	"github.com/google/uuid"
)

func UUID() string {
	return uuid.NewString()
}

func UUIDRaw() uuid.UUID {
	return uuid.New()
}
