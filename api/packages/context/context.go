package context

import (
	"backend/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/jinzhu/gorm"
)

type Context interface {
	Authenticated() bool
	RDB() *gorm.DB
	DynamoDB() *dynamo.DB
	Claim() *entity.Claim
}

type ctx struct {
	getRDB      func() *gorm.DB
	rdb         *gorm.DB
	getDynamoDB func() *dynamo.DB
	dynamoDB    *dynamo.DB
	claim       *entity.Claim
}

func New(c *gin.Context, getRDB func() *gorm.DB, getDynamoDB func() *dynamo.DB) Context {
	claim := entity.NewClaim(c)
	return &ctx{
		getRDB:      getRDB,
		getDynamoDB: getDynamoDB,
		claim:       &claim,
	}
}

func (c *ctx) RDB() *gorm.DB {
	if c.rdb != nil {
		return c.rdb
	}
	c.rdb = c.getRDB()
	return c.rdb
}

func (c *ctx) DynamoDB() *dynamo.DB {
	if c.dynamoDB != nil {
		return c.dynamoDB
	}
	c.dynamoDB = c.getDynamoDB()
	return c.dynamoDB
}

func (c *ctx) Authenticated() bool {
	return true
}

func (c *ctx) Claim() *entity.Claim {
	return c.claim
}
