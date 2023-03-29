/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"net/http"
	"testing"
)

func TestCert(t *testing.T) {
	cert := NewDefault("liaorenj@gmail.com", []string{"card.renj.io"})
	t.Log(cert.Ready())

	cert.SetChallengePath("/home")
	t.Log(cert.Ready())
	res, err := cert.Create()
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("cer: %s", res.Certificate)
	t.Logf("key: %s", res.PrivateKey)
}

func TestListen80(t *testing.T) {
	t.Log(canRunHTTP())
}

func TestListen(t *testing.T) {
	http.ListenAndServe(":80", nil)
}
