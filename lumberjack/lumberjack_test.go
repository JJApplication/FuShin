/*
   Create: 2023/8/19
   Project: FuShin
   Github: https://github.com/landers1037
   Copyright Renj
*/

package lumberjack

import (
	"log"
	"testing"
)

func TestNewRotate(t *testing.T) {
	r := NewEmptyRotate()
	log.SetOutput(r)

	log.Println("test log rotate")
}
