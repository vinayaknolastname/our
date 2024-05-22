package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
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
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Printf("host [%s]\n", config.GrpcPort)
	// fmt.Printf("Port [%d]\n", ConfigStruct.Port)
	// fmt.Printf("Enabled [%t]", ConfigStruct.Enabled)

	return
}
