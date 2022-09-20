package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// 自定义的payload, 可以嵌入标准Claims
type Payload struct {
	jwt.StandardClaims
	Sub  string `json:"sub"`
	User string `json:"user"`
}

const Secret = "jwt_test"

func main() {
	token := GenToken()
	fmt.Println(token)
	if ValidToken(token) {
		fmt.Println("valid jwt token")
	} else {
		fmt.Println("invalid jwt token")
	}

}

func GenToken() string {
	h := &Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	p := &Payload{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "lfn",
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
		Sub:  "test",
		User: "lfn",
	}

	hJson, _ := json.Marshal(h)
	pJson, _ := json.Marshal(p)
	hBase64URL := base64.URLEncoding.EncodeToString(hJson)
	pBase64URL := base64.URLEncoding.EncodeToString(pJson)

	// 计算hash之前剪除padding字符
	data := strings.TrimRight(hBase64URL, "=") + "." + strings.TrimRight(pBase64URL, "=")

	sig := SigBase64URLEncoded(data)

	token := data + "." + sig

	return token
}

func ValidToken(token string) bool {
	idx := strings.LastIndex(token, ".")
	data, sign := token[:idx], token[idx+1:]

	sigFromData := SigBase64URLEncoded(data)

	fmt.Printf("\n%-5s:%30s\n", "want", sign)
	fmt.Printf("%-5s:%3s\n", "got", sigFromData)

	return sigFromData == sign
}

func SigBase64URLEncoded(data string) string {
	ha := hmac.New(sha256.New, []byte(Secret))
	// Write Data to it
	ha.Write([]byte(data))
	sha := ha.Sum(nil)

	// base64URL编码
	rawSig := base64.URLEncoding.EncodeToString(sha)

	return strings.TrimRight(rawSig, "=")
}
