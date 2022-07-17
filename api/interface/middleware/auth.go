package middleware

import (
	"backend/config"
	"backend/infrastructure/database"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var issuer = fmt.Sprintf("https://cognito-idp.ap-northeast-1.amazonaws.com/%s", config.Env.AWS.USER_POOL_ID)

type Token struct {
	Payload  string `dynamo:"payload,hash"`
	Disabled bool   `dynamo:"disabled"`
	Ttl      int    `dynamo:"ttl"`
}

func Auth(c *gin.Context) {
	authorizationValue := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(authorizationValue, " ")[1]

	claim, err := auth(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if claim == nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	ddb := database.GetDynamoDB()
	table := ddb.Table("Token")
	var token Token

	payload := strings.Split(tokenString, ".")[1]
	err = table.Get("payload", payload).One(&token)
	if err != nil {
		err = table.Put(&Token{Payload: payload, Disabled: false, Ttl: claim.Exp}).Run()
		if err != nil {
			panic(err)
		}
		c.Next()
		return
	}

	if token.Disabled {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
	return

}

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

func decodeJwt(token string) (*Header, *Claim, error) {
	tokenSections := strings.Split(token, ".")
	if len(tokenSections) < 2 {
		panic("requested token is invalid")
	}
	headerSection := tokenSections[0]
	payloadSection := tokenSections[1]

	decodedHeader, _ := base64.RawURLEncoding.DecodeString(headerSection)
	decodedPayload, _ := base64.RawURLEncoding.DecodeString(payloadSection)

	header := Header{}
	claim := Claim{}

	json.Unmarshal(decodedHeader, &header)
	json.Unmarshal(decodedPayload, &claim)

	return &header, &claim, nil
}

func verify(tokenString string, key Key) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		key := convertKey(key)
		return key, nil
	})
	return token, err
}

func verifyClaim(claim Claim) (ok bool) {
	ok = true
	// NOTE: 有効期限チェック
	currentSec := time.Now().Unix()
	if currentSec > int64(claim.Exp) {
		fmt.Println("token is expired")
	}
	// NOTE: JWT発行者チェック
	if claim.Iss != issuer {
		fmt.Println("invalid iss")
	}
	// 使用トークンチェック
	if claim.TokenUse != "access" {
		fmt.Println("invalid token use")
	}
	return ok
}

func convertKey(key Key) *rsa.PublicKey {
	decodedE, err := base64.RawStdEncoding.DecodeString(key.E)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	publicKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		panic(err)
	}
	publicKey.N.SetBytes(decodedN)
	return publicKey
}

func auth(tokenString string) (*Claim, error) {
	header, claim, err := decodeJwt(tokenString)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/.well-known/jwks.json", issuer)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jwk := Jwk{}
	byteArray, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(byteArray, &jwk)

	key := Key{}
	for _, v := range jwk.Keys {
		if v.Kid == header.Kid {
			key = v
		}
	}

	token, err := verify(tokenString, key)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	isValidClaim := verifyClaim(*claim)
	if !isValidClaim {
		return nil, errors.New("in valid claim")
	}

	return claim, nil
}

/**
@see https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
@see https://aws.amazon.com/jp/premiumsupport/knowledge-center/decode-verify-cognito-json-token/
@see https://qiita.com/aioa/items/8bc1eb0d021745f8ea85
*/
