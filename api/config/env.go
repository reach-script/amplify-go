package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Env EnvType

type EnvType struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
		SSL_MODE string
	}
	API struct {
		Port string
	}
	AWS struct {
		USER_POOL_ID      string
		REGION            string
		DYNAMO_ENDPOINT   string
		ACCESS_KEY_ID     string
		SECRET_ACCESS_KEY string
	}
}

func envExists(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		return false
	}
	return true
}

func init() {
	dotenvPath := ".env"
	if dotenvPathEnv := os.Getenv("DOTENV_PATH"); dotenvPathEnv != "" {
		dotenvPath = dotenvPathEnv
	}
	if envExists(dotenvPath) {
		err := godotenv.Load(dotenvPath)
		if err != nil {
			panic(err)
		}
	}

	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", Env)
}
