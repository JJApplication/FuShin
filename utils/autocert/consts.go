/*
Create: 2023/3/28
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

const (
	CertPath = "fushin-certs" // 存储ssl证书的路径
	Renew    = 60 * 60 * 24 * 30
)

const (
	RSA2048 string = "RSA-2048"
	RSA4096 string = "RSA-4096"
	RSA8192 string = "RSA-8192"
	EC256   string = "ECDSA-256"
	EC384   string = "ECDSA-384"
)

const (
	errEmptyDomains = "domains are empty"
	errEmptyEmail   = "email is empty"
	errNoCert       = "certDir not set"
	errNoChallenge  = "challengeDir not set"
	errHttpBusy     = "http port 80 is already in use"
	errCertNoEmpty  = "certDir not empty"
)
