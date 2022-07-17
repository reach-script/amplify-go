package database

import (
	"fmt"

	"backend/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var rdb *gorm.DB
var ddb *dynamo.DB

func init() {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Env.DB.User,
		config.Env.DB.Password,
		config.Env.DB.Host,
		config.Env.DB.Port,
		config.Env.DB.Name,
		config.Env.DB.SSL_MODE)

	var err error
	rdb, err = gorm.Open("postgres", connectionStr)
	rdb.LogMode(true)
	if err != nil {
		panic(err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Env.AWS.REGION),
		Endpoint:    aws.String(config.Env.AWS.DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials(config.Env.AWS.ACCESS_KEY_ID, config.Env.AWS.SECRET_ACCESS_KEY, ""),
	})
	if err != nil {
		panic(err)
	}

	ddb = dynamo.New(sess)
}

func GetRDB() *gorm.DB {
	return rdb
}

func GetDynamoDB() *dynamo.DB {
	return ddb
}
