package context

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Context interface {
	Authenticated() bool
	DB() *gorm.DB
}

type ctx struct {
	getDB func() *gorm.DB
	db    *gorm.DB
}

func New(c *gin.Context, getDB func() *gorm.DB) Context {
	return &ctx{
		getDB: getDB,
	}
}

func (c *ctx) DB() *gorm.DB {
	if c.db != nil {
		return c.db
	}
	c.db = c.getDB()
	return c.db
}

func (c *ctx) Authenticated() bool {
	return true
}
