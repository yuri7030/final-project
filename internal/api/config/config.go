package config

import (
	"github.com/spf13/viper"
	"github.com/yuri7030/final-project/internal/constants"
)

// AppConfig represents the structure of the application configuration.
type AppConfig struct {
	Environment string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	JwtSecret   string
}

// LoadConfig loads the application configuration from a config file.
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("development")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.SetDefault("environment", "development")
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "12345678")
	viper.SetDefault("db_name", "ProductDB")
	viper.SetDefault("jwt_secret", constants.JwtSecretKey)

	appConfig := &AppConfig{
		Environment: viper.GetString("environment"),
		DBUser:      viper.GetString("db_user"),
		DBPassword:  viper.GetString("db_password"),
		DBName:      viper.GetString("db_name"),
		DBPort:      viper.GetString("db_port"),
		JwtSecret:   viper.GetString("jwt_secret"),
	}

	return appConfig, nil
}

func GetEnvironment() string {
	return viper.GetString("environment")
}

func GetDbUser() string {
	return viper.GetString("db_user")
}

func GetDbPassword() string {
	return viper.GetString("db_password")
}

func GetDbName() string {
	return viper.GetString("db_name")
}

func GetJwtSecret() string {
	return viper.GetString("jwt_secret")
}

func GetDbPort() string {
	return viper.GetString("db_port")
}
