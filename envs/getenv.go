package envs

import (
	"log"

	"github.com/spf13/viper"
)

func Getenv(envFile string) string {

	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Default().Fatalln("Could not load environment variables", err)
	}

	return viper.GetString(envFile)
}
