/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"fmt"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
)

// NewClient returns a new Lets Encrypt client
func NewClient(account *Cert) (*lego.Client, error) {
	cfg := lego.NewConfig(account)
	switch account.KeyType {
	case RSA2048:
		cfg.Certificate.KeyType = certcrypto.RSA2048
		return lego.NewClient(cfg)
	case RSA4096:
		cfg.Certificate.KeyType = certcrypto.RSA4096
		return lego.NewClient(cfg)
	case RSA8192:
		cfg.Certificate.KeyType = certcrypto.RSA8192
		return lego.NewClient(cfg)
	case EC256:
		cfg.Certificate.KeyType = certcrypto.EC256
		return lego.NewClient(cfg)
	case EC384:
		cfg.Certificate.KeyType = certcrypto.EC384
		return lego.NewClient(cfg)
	default:
		return nil, fmt.Errorf("invalid private key type: %s", account.KeyType)
	}
}
