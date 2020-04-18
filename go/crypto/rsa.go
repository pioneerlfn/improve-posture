package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// 公钥加密
func encrypt(raw string, publicKey rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(raw),
		nil,
	)
	return encryptedBytes, err
}

// 私钥解密
func decrypt(private *rsa.PrivateKey, data []byte) ([]byte, error) {
	decryptedBytes, err := private.Decrypt(
		nil,
		data,
		&rsa.OAEPOptions{Hash: crypto.SHA256},
	)
	return decryptedBytes, err
}

// ----------------

// 私钥签名(证明身份)
func sig(msg []byte, private *rsa.PrivateKey) ([]byte, error) {
	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(
		rand.Reader,
		private,
		crypto.SHA256,
		msgHashSum,
		nil)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// 公钥验证身份
func verifySig(msg, signature []byte, public rsa.PublicKey) error {
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		return err
	}
	sum := h.Sum(nil)
	err = rsa.VerifyPSS(&public, crypto.SHA256, sum, signature, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	message := "super secret message"
	// 加密
	enc, err := encrypt(message, privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密信息是:\n0x%x\n", enc)

	// 解密
	dec, err := decrypt(privateKey, enc)
	if err != nil {
		panic(err)
	}

	if message != string(dec) {
		panic("解密信息不等于原信息")
	}
	fmt.Printf("\n解密成功: ")
	fmt.Println(string(dec))

	// 签名
	signature, err := sig([]byte(message), privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n===========\n签名是:\n0x%x\n", signature)

	// 验证
	err = verifySig([]byte(message), signature, privateKey.PublicKey)
	if err != nil {
		fmt.Printf("\n签名验证失败")
		return
	}
	fmt.Println("\n签名验证成功")
}

// 参考[Implementing RSA Encryption and Signing in Golang (With Examples)](https://www.sohamkamani.com/golang/rsa-encryption/)
