/*
   Create: 2023/8/19
   Project: FuShin
   Github: https://github.com/landers1037
   Copyright Renj
*/

package lumberjack

import "time"

// RegisterRotate create a new go routine
// unit seconds
func RegisterRotate(r Rotate, duration int) {
	t := time.NewTicker(time.Duration(duration) * time.Second)
	go func() {
		for {
			<-t.C
			r.Rotate()
		}
	}()
}
