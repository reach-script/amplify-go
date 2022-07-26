package middleware

import (
	"backend/config"
	"backend/domain/entity"
	"backend/infrastructure/database"
	"backend/packages/utils"
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
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

var issuer = fmt.Sprintf("https://cognito-idp.ap-northeast-1.amazonaws.com/%s", config.Env.AWS.USER_POOL_ID)

func Auth(c *gin.Context) {
	claim, jwt, err := auth(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if claim == nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	db := database.GetDynamoDB()
	table := db.Table("Session")
	var auth entity.Auth

	err = table.Get("key1", claim.Sub).Range("key2", dynamo.Equal, jwt.Payload).One(&auth)
	if err != nil {
		item := entity.Auth{Key1: claim.Sub, Key2: jwt.Payload, Payload: jwt.Payload, Disabled: false, Ttl: claim.Exp}
		err = table.Put(&item).Run()
		if err != nil {
			panic(err)
		}
		c.Next()
		return
	}

	if auth.Disabled {
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

func verify(tokenString string, key Key) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		key := convertKey(key)
		return key, nil
	})
	return token, err
}

func verifyClaim(claim entity.Claim) (ok bool) {
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

func auth(c *gin.Context) (*entity.Claim, *utils.Jwt, error) {
	tokenString := utils.GetJwt(c)
	jwt := utils.NewJwt(tokenString)
	header, claim, err := utils.DecodeJwt(jwt)
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("%s/.well-known/jwks.json", issuer)

	response, err := http.Get(url)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}
	if !token.Valid {
		return nil, nil, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, nil, err
	}

	isValidClaim := verifyClaim(*claim)
	if !isValidClaim {
		return nil, nil, errors.New("in valid claim")
	}

	return claim, &jwt, nil
}

/**
@see https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
@see https://aws.amazon.com/jp/premiumsupport/knowledge-center/decode-verify-cognito-json-token/
@see https://qiita.com/aioa/items/8bc1eb0d021745f8ea85
*/
