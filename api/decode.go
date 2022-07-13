package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type publicKey struct {
	Keys []key `json:"keys"`
}
type key struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

type payload struct {
	Sub       string `json:"sub"`
	Iss       string `json:"iss"`
	ClientID  string `json:"client_id"`
	OriginJti string `json:"origin_jti"`
	EventID   string `json:"event_id"`
	TokenUse  string `json:"token_use"`
	Scope     string `json:"scope"`
	AuthTime  int    `json:"auth_time"`
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
	Jti       string `json:"jti"`
	Username  string `json:"username"`
}

func decodeJwt(token string) (string, error) {
	divided := strings.Split(token, ".")
	temp := strings.Replace(divided[1], "-", "+", -1)
	placed := strings.Replace(temp, "_", "/", -1)

	llx := len(placed)
	nnx := ((4 - llx%4) % 4)
	ssx := strings.Repeat("=", nnx)
	joined := strings.Join([]string{placed, ssx}, "")
	str, err := b64.StdEncoding.DecodeString(joined)
	if err != nil {
		return "", err
	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	decode, _ := b64.URLEncoding.DecodeString(uEnc)

	return string(decode), nil
}

func main() {
	buff, _ := ioutil.ReadFile("token01.txt")
	decoded, err := decodeJwt(string(buff))

	if err != nil {
		panic(err)
	}

	payloadObj := payload{}
	json.Unmarshal([]byte(decoded), &payloadObj)

	fmt.Println(payloadObj.Iss)

	url := fmt.Sprintf("%s/.well-known/jwks.json", payloadObj.Iss)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	obj := publicKey{}
	byteArray, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(byteArray, &obj)

	for _, v := range obj.Keys {
		fmt.Println(v)
	}
}

/**
{
	"sub":"11ad928e-a71a-46ec-840f-0572da32aa97",
	"iss":"https:\/\/cognito-idp.ap-northeast-1.amazonaws.com\/ap-northeast-1_Oxzc2wtHu",
	"client_id":"60b9pj6tmet4gsgn0kdf6quv3b",
	"origin_jti":"83dc8b35-93d4-4004-acf3-eac2aaf24a28",
	"event_id":"939455ef-feb8-4a42-9048-e23299cb9d66",
	"token_use":"access",
	"scope":"aws.cognito.signin.user.admin",
	"auth_time":1657631298,
	"exp":1657634898,
	"iat":1657631298,
	"jti":"81d1fff5-e906-4457-9a7c-766a78512784",
	"username":"11ad928e-a71a-46ec-840f-0572da32aa97"
}
*/

/**
@see https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
@see https://aws.amazon.com/jp/premiumsupport/knowledge-center/decode-verify-cognito-json-token/
@see https://qiita.com/aioa/items/8bc1eb0d021745f8ea85
*/
