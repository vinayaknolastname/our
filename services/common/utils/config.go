package utils

import (
	"log"

	"github.com/dvln/viper"
)

type Config struct {
	GrpcPort string `mapstructure:"GRPC_SERVER_ADDRESS"`
	DBConfig *DBConfig
}

type DBConfig struct {
	DBHost     string `mapstructure:"DB_Host"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBNAME     string `mapstructure:"DB_NAME"`
	DBUSER     string `mapstructure:"DB_USER"`
	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Println("cofig load error  %err ", err)

	}

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}