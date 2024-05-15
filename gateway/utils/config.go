package utils

type Config struct {
	GatewayPort int32
	Cloudinary  string
}

func GetConfig() *Config {
	return &Config{
		GatewayPort: 5000,
		Cloudinary:  "ddd",
	}
}

// type DBConfig struct {
// 	DBHost     string `mapstructure:"DB_Host"`
// 	DBPort     int    `mapstructure:"DB_PORT"`
// 	DBNAME     string `mapstructure:"DB_NAME"`
// 	DBUSER     string `mapstructure:"DB_USER"`
// 	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
// }

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)

// 	viper.SetConfigName("app")
// 	viper.SetConfigType("yaml")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()

// 	if err != nil {
// 		log.Println("cofig load error  %err ", err)

// 	}

// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }
