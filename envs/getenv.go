package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(envFile string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Cannot read your environment variable", err)
	}
	return os.Getenv(envFile)
}
