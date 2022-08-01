package entity

import (
	"backend/config"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

type header struct {
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type Jwt struct {
	Header        string
	Payload       string
	Signature     string
	DecodedHeader *header
	Claim         Claim
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

	decodedHeader, _ := base64.RawURLEncoding.DecodeString(jwt.Header)
	decodedPayload, _ := base64.RawURLEncoding.DecodeString(jwt.Payload)

	header := header{}
	claim := Claim{}

	json.Unmarshal(decodedHeader, &header)
	json.Unmarshal(decodedPayload, &claim)

	jwt.DecodedHeader = &header
	jwt.Claim = claim

	return jwt
}

func (jwt *Jwt) Validate(iss string) bool {
	claim := jwt.Claim

	if claim.Exp < int(time.Now().Unix()) {
		return false
	}
	if claim.ClientID != config.Env.AWS.Cognito.APP_CLIENT_ID {
		return false
	}
	if claim.Iss != iss {
		return false
	}
	if claim.TokenUse != "access" {
		return false
	}
	return true
}
