package utils

import (
	"backend/domain/entity"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type Jwk struct {
	Keys []Key
}

type Key struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

type Jwt struct {
	Header    string
	Payload   string
	Signature string
}

func NewJwt(token string) Jwt {
	tokenSections := strings.Split(token, ".")
	if len(tokenSections) < 2 {
		panic("requested token is invalid")
	}

	jwt := Jwt{
		Header:    tokenSections[0],
		Payload:   tokenSections[1],
		Signature: tokenSections[2],
	}

	return jwt
}

func DecodeJwt(jwt Jwt) (*Header, *entity.Claim, error) {
	decodedHeader, _ := base64.RawURLEncoding.DecodeString(jwt.Header)
	decodedPayload, _ := base64.RawURLEncoding.DecodeString(jwt.Payload)

	header := Header{}
	claim := entity.Claim{}

	json.Unmarshal(decodedHeader, &header)
	json.Unmarshal(decodedPayload, &claim)

	return &header, &claim, nil
}

func GetJwt(c *gin.Context) string {
	authorizationValue := c.Request.Header.Get("Authorization")
	token := strings.Split(authorizationValue, " ")[1]

	return token
}
