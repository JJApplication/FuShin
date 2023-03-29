/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/http/webroot"
	"github.com/go-acme/lego/v4/registration"
)

// 仅支持httpProvider

// RegProvider 1.注册质询
func RegProvider(c *lego.Client, challenge string) error {
	var err error
	ps, err := webroot.NewHTTPProvider(challenge)
	if err != nil {
		return err
	}
	err = c.Challenge.SetHTTP01Provider(ps)
	if err != nil {
		return err
	}
	return err
}

// RegAccount 2.注册账户
func RegAccount(c *lego.Client) (*registration.Resource, error) {
	reg, err := c.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return reg, err
	}
	return reg, nil
}
