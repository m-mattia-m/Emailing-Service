package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	env, err := os.LookupEnv(key) // GetEnv() does not work
	if !err || env == "" {
		return env, errors.New("getenv: environment variable empty")
	}
	return env, nil
}

func init() {
	log.Println("[INIT MAIN]: Status -> Main-init is loaded")
	godotenv.Load()
}
