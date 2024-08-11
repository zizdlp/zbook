package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment            string        `mapstructure:"ENVIRONMENT"`
	DBSource               string        `mapstructure:"DB_SOURCE"`
	MigrationURL           string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress      string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	HOMEADDRESS            string        `mapstructure:"HOME_ADDRESS"`
	GRPCServerAddress      string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	WEBSOCKETServerAddress string        `mapstructure:"WEBSOCKET_SERVER_ADDRESS"`
	RedisAddress           string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey      string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	REQUIRE_EMAIL_VERIFY   bool          `mapstructure:"REQUIRE_EMAIL_VERIFY"`
	EmailSenderName        string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress     string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword    string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	SmtpAuthAddress        string        `mapstructure:"SMTP_AUTH_ADDR"`
	SmtpServerAddress      string        `mapstructure:"SMTP_SERVER_ADDR"`
	MINIOADDR              string        `mapstructure:"MINIO_ADDR"`
	MINIOROOTPASSWORD      string        `mapstructure:"MINIO_ROOT_PASSWORD"`
	MINIOROOTUSER          string        `mapstructure:"MINIO_ROOT_USER"`
	TIMEZONE               string        `mapstructure:"TIME_ZONE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
