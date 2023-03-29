/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
)

// ObtainDomains 3.注册域名
func ObtainDomains(c *lego.Client, domains []string) (*certificate.Resource, error) {
	req := certificate.ObtainRequest{Domains: domains, Bundle: true}
	return c.Certificate.Obtain(req)
}
