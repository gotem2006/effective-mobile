package config

import "github.com/spf13/viper"



type Config struct
{
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBName string `mapstructure:"DB_NAME"`
	DBUser string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	APIUrl string `mapstructure:"API_URL"`
}


func LoadConfig(path string) (Config, error){
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil{
		return config, err
	}

	err := viper.Unmarshal(&config)

	return config, err
}