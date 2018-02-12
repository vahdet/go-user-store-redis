package app

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

// Config stores the application-wide configurations
var validate *validator.Validate

var Config appConfig

type appConfig struct {
	GrpcPort string `mapstructure:"grpc.port"`
	ErrorFile string `mapstructure:"error_file"`
	DataStoreUrl string `mapstructure:"datastore_url" validate:"required"`
	JWTSigningMethod string `mapstructure:"jwt.signing_method"`
	JWTSigningKey string `mapstructure:"jwt.signing_key"`
	JWTVerificationKey string `mapstructure:"jwt.verification_key"`
}

// Deprecated read from env variables instead of files
// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	// File name and extension
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	// watch the changes in the config file and prompt a message is there are any
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	// Environmennt
	// v.SetEnvPrefix("restful")
	v.AutomaticEnv()
	// Defaults
	v.SetDefault("error_file", "resources/errors.yaml")
	v.SetDefault("server_port", 8080)
	// Add config paths given as the parameter
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	// validation
	validate = validator.New()

	return validate.Struct(&Config)
}
