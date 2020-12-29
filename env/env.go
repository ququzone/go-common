package env

import (
	"log"
	"os"
)

func GetNonEmpty(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalf("%s is not defined in env\n", key)
	}
	return str
}
