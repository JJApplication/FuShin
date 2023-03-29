/*
Create: 2023/3/28
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"crypto"
	"errors"
	"io/ioutil"
	"os"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/registration"
)

// NewCert 根据缓存目录创建证书生成器
func NewCert(certPath, email string, domains []string) Cert {
	SetCertCacheDir(certPath)

	return Cert{
		Email:     email,
		certPath:  certPath,
		domains:   domains,
		renewTime: 0,
	}
}

func NewDefault(email string, domains []string) Cert {
	SetCertCacheDir(CertPath)

	return Cert{
		Email:     email,
		certPath:  CertPath,
		domains:   domains,
		renewTime: 0,
	}
}

type Cert struct {
	Email        string
	Registration *registration.Resource
	KeyType      string
	Challenge    string // 质询文件路径 默认拼接$PATH/.well-known/acme-challenge/
	key          crypto.PrivateKey
	certPath     string
	domains      []string
	renewTime    int // 单位s 默认30天 60 * 60 * 24 * 30
}

func (c *Cert) SetCertPath(p string) {
	c.certPath = p
}

func (c *Cert) SetChallengePath(p string) {
	c.Challenge = p
}

func (c *Cert) SetKeyType(t string) {
	c.KeyType = t
}

func (c *Cert) AddDomains(ds []string) {
	c.domains = ds
}

func (c *Cert) SetRenew(t int) {
	if t <= 0 {
		c.renewTime = Renew
	} else {
		c.renewTime = t
	}
}

func (c *Cert) GetCertPath() string {
	return c.certPath
}

func (c *Cert) GetDomains() []string {
	return c.domains
}

// Ready 是否准备就绪
func (c *Cert) Ready() error {
	if len(c.domains) <= 0 {
		return errors.New(errEmptyDomains)
	}
	if c.Email == "" {
		return errors.New(errEmptyEmail)
	}
	if c.certPath == "" {
		return errors.New(errNoCert)
	}
	if c.Challenge == "" {
		return errors.New(errNoChallenge)
	}
	if !canRunHTTP() {
		return errors.New(errHttpBusy)
	}
	return nil
}

// Create 返回生成的证书列表
func (c *Cert) Create() (*certificate.Resource, error) {
	if c.Ready() != nil {
		return nil, c.Ready()
	}
	dir, err := ioutil.ReadDir(c.certPath)
	if !os.IsNotExist(err) {
		if len(dir) <= 0 {
			return nil, errors.New(errCertNoEmpty)
		}
	}
	return c.start()
}

// Run 作为自动更新的服务端启动
func (c *Cert) Run() (err error) {
	if c.Ready() != nil {
		return c.Ready()
	}
	AutoCert(c, err)
	return err
}

// RunAndStop 仅用于下载申请证书 在占用完毕80端口后会释放
func (c *Cert) RunAndStop() error {
	if c.Ready() != nil {
		return c.Ready()
	}
	dir, err := ioutil.ReadDir(c.certPath)
	if !os.IsNotExist(err) {
		if len(dir) <= 0 {
			return errors.New(errCertNoEmpty)
		}
	}
	res, err := c.start()
	if err != nil {
		return err
	}
	return SaveCert(c.certPath, res)
}

/* Methods implementing the lego.User interface*/

// GetEmail returns the email address for the account
func (c *Cert) GetEmail() string {
	return c.Email
}

// GetPrivateKey returns the private RSA account key.
func (c *Cert) GetPrivateKey() crypto.PrivateKey {
	return c.key
}

// GetRegistration returns the server registration
func (c *Cert) GetRegistration() *registration.Resource {
	return c.Registration
}

func (c *Cert) AddRegistration(reg *registration.Resource) {
	c.Registration = reg
}

func (c *Cert) start() (*certificate.Resource, error) {
	if c.KeyType == "" {
		c.KeyType = RSA2048
	}
	priv, err := GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	c.key = priv

	cli, err := NewClient(c)
	if err != nil {
		return nil, err
	}

	if err = RegProvider(cli, c.Challenge); err != nil {
		return nil, err
	}

	reg, err := RegAccount(cli)
	if err != nil {
		return nil, err
	}

	c.AddRegistration(reg)

	certs, err := ObtainDomains(cli, c.domains)

	return certs, err
}
