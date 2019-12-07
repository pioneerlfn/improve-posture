# JWT

`jwt`看[官方文档](https://jwt.io/introduction/) 就够了。下面用`Go`实践一下:

## 生成token

```Go

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// 自定义的payload, 可以嵌入标准Claims
type Payload struct {
	Sub  string `json:"sub"`
	User string `json:"user"`
}

const Secret = "jwt_test"

func main() {
	h := &Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	p := &Payload{
		Sub:  "test",
		User: "lfn",
	}

	hJson, _ := json.Marshal(h)
	pJson, _ := json.Marshal(p)
	hBase64URL := base64.URLEncoding.EncodeToString(hJson)
	pBase64URL := base64.URLEncoding.EncodeToString(pJson)

	data := hBase64URL + "." + pBase64URL


	// Create a new HMAC by defining the hash type and the key (as byte array)
	ha := hmac.New(sha256.New, []byte(Secret))

	// Write Data to it
	ha.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := ha.Sum(nil)

	// base64URL编码
	rawSig := base64.URLEncoding.EncodeToString(sha)

	// 去掉右侧padding符
	sig := strings.TrimRight(rawSig, "=")

	jwt := data + "." + sig
    fmt.Println(jwt) 
    // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwidXNlciI6ImxmbiJ9.UarPvSxHs3ywhTvXHs2Kc4GRwcEWqdah8tdFvY0nAWY
}

```

> 需要注意的是，计算出签名之后，仍然需要用Base64URL编码，并且需要祛除右侧的padding符"=".

