package middleware

import (
	"backend/config"
	"backend/domain/entity"
	"backend/infrastructure/database"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

type jwk struct {
	Keys []key
}

type key struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

var issuer = fmt.Sprintf("https://cognito-idp.ap-northeast-1.amazonaws.com/%s", config.Env.AWS.Cognito.USER_POOL_ID)

func Auth(c *gin.Context) {
	authorizationValue := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(authorizationValue, " ")[1]
	jwt := entity.NewJwt(tokenString)

	if ok := jwt.Validate(issuer); !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	key, err := getKey(jwt.DecodedHeader.Kid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ok, err := verifyToken(tokenString, key)
	if !ok || err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	db := database.GetDynamoDB()
	table := db.Table("Session")
	var auth entity.Auth

	err = table.Get("key1", jwt.Claim.Sub).Range("key2", dynamo.Equal, jwt.Payload).One(&auth)
	if err != nil {
		item := entity.Auth{Sub: jwt.Claim.Sub, Payload: jwt.Payload, Disabled: false, Ttl: jwt.Claim.Exp}
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

func verifyToken(tokenString string, key *key) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		key := convertKey(key)
		return key, nil
	})

	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, err
	}
	if err := token.Claims.Valid(); err != nil {
		return false, err
	}

	return true, err
}

func convertKey(key *key) *rsa.PublicKey {
	decodedE, err := base64.RawStdEncoding.DecodeString(key.E)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		data := make([]byte, 4)
		copy(data[4-len(decodedE):], decodedE)
		decodedE = data
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

func getKey(kid string) (*key, error) {
	url := fmt.Sprintf("%s/.well-known/jwks.json", issuer)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jwk := jwk{}
	byteArray, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(byteArray, &jwk)

	var key *key = nil
	for _, v := range jwk.Keys {
		if v.Kid == kid {
			key = &v
		}
	}

	return key, nil
}

/**
@see https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
@see https://aws.amazon.com/jp/premiumsupport/knowledge-center/decode-verify-cognito-json-token/
@see https://qiita.com/aioa/items/8bc1eb0d021745f8ea85
*/
