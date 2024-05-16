package config

import (
	"github.com/spf13/viper"
)

var (
	OTELEndpoint string
	ServiceName  string
	ServicePort  string
	IsDebugMode  bool
)

func Init() {
	ServiceName = "learning-o11y"

	viper.AutomaticEnv()

	viper.SetDefault("DEBUG", true)
	IsDebugMode = viper.GetBool("DEBUG")

	viper.SetDefault("OTEL_ENDPOINT", "localhost:4317")
	OTELEndpoint = viper.GetString("OTEL_ENDPOINT")

	viper.SetDefault("SERVICE_PORT", "8080")
	ServicePort = viper.GetString("SERVICE_PORT")
}
