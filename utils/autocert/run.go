/*
Create: 2023/3/28
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"context"
	"net/http"
)

func HttpRun() error {
	return http.ListenAndServe(":http", nil)
}

func HttpRunForOnce() error {
	return http.ListenAndServe(":http", nil)
}

func HttpRunWithCtx(ctx context.Context) {
	s := http.Server{
		Addr:    ":http",
		Handler: nil,
	}
	select {
	case <-ctx.Done():
		s.Close()
	default:
		s.ListenAndServe()
	}
}
