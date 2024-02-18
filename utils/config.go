package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type envConfigs struct {
	ETH_rpc string `mapstructure:"ETH_RPC"`
}

var EnvConfigs *envConfigs

// We will call this in main.go to load the env variables
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	// bind env variables. This will be used on production where we don't use the .env file
	// but rather load env from the environment variable
	viper.BindEnv("ETH_rpc", "ETH_RPC")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Load local .env in dev mod only
	if os.Getenv("ENV") == "development" {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error reading env file", err)
		}
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
