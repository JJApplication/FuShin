/*
Create: 2022/8/20
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package errors
package errors

import (
	"github.com/google/uuid"
)

func Panic(s string) {
	panic(FushinError{id: uuid.New().String(), message: s})
}
