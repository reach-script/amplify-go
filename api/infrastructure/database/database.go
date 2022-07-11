package database

import (
	"fmt"

	"backend/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func init() {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Env.DB.User,
		config.Env.DB.Password,
		config.Env.DB.Host,
		config.Env.DB.Port,
		config.Env.DB.Name,
		config.Env.DB.SSL_MODE)

	var err error
	db, err = gorm.Open("postgres", connectionStr)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	return db
}
