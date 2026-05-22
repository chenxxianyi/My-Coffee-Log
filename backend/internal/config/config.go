package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	AppPort        string `mapstructure:"APP_PORT"`
	MYSQLHost      string `mapstructure:"MYSQL_HOST"`
	MYSQLPort      string `mapstructure:"MYSQL_PORT"`
	MYSQLUser      string `mapstructure:"MYSQL_USER"`
	MYSQLPassword  string `mapstructure:"MYSQL_PASSWORD"`
	MYSQLDatabase  string `mapstructure:"MYSQL_DATABASE"`
	REDISHost      string `mapstructure:"REDIS_HOST"`
	REDISPort      string `mapstructure:"REDIS_PORT"`
	REDISPassword  string `mapstructure:"REDIS_PASSWORD"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JWTExpireHours int    `mapstructure:"JWT_EXPIRE_HOURS"`
	OpenAIAPIKey   string `mapstructure:"OPENAI_API_KEY"`
	OpenAIBaseURL  string `mapstructure:"OPENAI_BASE_URL"`
	OpenAIModel    string `mapstructure:"OPENAI_MODEL"`
	AIEnabled      bool   `mapstructure:"AI_ENABLED"`
	OpenAITimeout  int    `mapstructure:"OPENAI_REQUEST_TIMEOUT_SECONDS"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Failed to read config: %v", err)
		}
	}

	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.SetDefault("MYSQL_USER", "root")
	viper.SetDefault("MYSQL_PASSWORD", "123456")
	viper.SetDefault("MYSQL_DATABASE", "my_coffee_log")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("JWT_SECRET", "please_change_me")
	viper.SetDefault("JWT_EXPIRE_HOURS", 168)
	viper.SetDefault("AI_ENABLED", false)
	viper.SetDefault("OPENAI_REQUEST_TIMEOUT_SECONDS", 5)

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	if AppConfig.JWTSecret == "" || AppConfig.JWTSecret == "please_change_me" {
		if AppConfig.AppEnv == "production" {
			log.Fatalf("JWT_SECRET must be set to a secure value in production")
		}
		log.Println("WARNING: using default JWT secret, set JWT_SECRET env var for security")
	}
}
