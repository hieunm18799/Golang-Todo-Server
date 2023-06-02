package configs

import "github.com/spf13/viper"

type Config struct {
	Host       string `mapstructure:"POSTGRES_HOST"`
	User       string `mapstructure:"POSTGRES_USER"`
	Password   string `mapstructure:"POSTGRES_PASSWORD"`
	DB_Name    string `mapstructure:"POSTGRES_DB_NAME"`
	PSQL_Port  string `mapstructure:"POSTGRES_PORT"`
	ServerPort string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
