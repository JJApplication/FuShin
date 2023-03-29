/*
Create: 2023/3/30
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

package autocert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// GeneratePrivateKey 生成P256的私钥
func GeneratePrivateKey() (crypto.PrivateKey, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return pk, err
}

func GeneratePrivateKeyCustom(curve elliptic.Curve) (crypto.PrivateKey, error) {
	pk, err := ecdsa.GenerateKey(curve, rand.Reader)
	return pk, err
}
