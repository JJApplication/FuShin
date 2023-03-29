/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-acme/lego/v4/certificate"
)

const (
	PrivateKey  = "private.key"   // 私钥
	Certificate = "fullchain.pem" // 证书
	CA          = "ca.crt"        // CA
	CSR         = "csr"           // 全量csr
)

// SaveCert 保存cert文件
// 默认会保存到$CertRoot下
func SaveCert(p string, cert *certificate.Resource) error {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		if e := os.MkdirAll(p, 0755); e != nil {
			return err
		}
	}
	// write priv
	err := writeFile(p, PrivateKey, cert.PrivateKey)
	if err != nil {
		return err
	}
	err = writeFile(p, Certificate, cert.Certificate)
	if err != nil {
		return err
	}
	err = writeFile(p, CA, cert.IssuerCertificate)
	if err != nil {
		return err
	}
	err = writeFile(p, CSR, cert.CSR)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(p, f string, data []byte) error {
	fp := filepath.Join(p, f)
	return ioutil.WriteFile(fp, data, 0644)
}
