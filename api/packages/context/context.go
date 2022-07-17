package context

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/jinzhu/gorm"
)

type Context interface {
	Authenticated() bool
	RDB() *gorm.DB
	DynamoDB() *dynamo.DB
}

type ctx struct {
	getRDB      func() *gorm.DB
	rdb         *gorm.DB
	getDynamoDB func() *dynamo.DB
	dynamoDB    *dynamo.DB
}

func New(c *gin.Context, getRDB func() *gorm.DB, getDynamoDB func() *dynamo.DB) Context {
	return &ctx{
		getRDB:      getRDB,
		getDynamoDB: getDynamoDB,
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
