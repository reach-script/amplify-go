package entity

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

type Claim struct {
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

func NewClaim(c *gin.Context) Claim {
	authorizationValue := c.Request.Header.Get("Authorization")
	token := strings.Split(authorizationValue, " ")[1]
	claim := decode(token)
	return claim
}

func decode(token string) Claim {
	tokenSections := strings.Split(token, ".")
	if len(tokenSections) < 2 {
		panic("requested token is invalid")
	}
	payloadSection := tokenSections[1]

	decodedPayload, _ := base64.RawURLEncoding.DecodeString(payloadSection)

	claim := Claim{}

	json.Unmarshal(decodedPayload, &claim)

	return claim
}
