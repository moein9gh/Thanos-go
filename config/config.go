package config

import (
	"bytes"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/thanos-go/log"

	"github.com/renstrom/shortuuid"
	"github.com/spf13/viper"
)

const (
	// ProductionMode indicates mode is release.
	ProductionMode = "production"
	// TestingMode indicates mode is test.
	TestingMode = "testing"
	// DebuggingMode indicates mode is debug.
	DebuggingMode = "debugging"

	// KeyDeviceClaim is used to store the device data in the context
	KeyDeviceClaim = "device_claims"
)

// IgnoredPathsInLogs all requests to these endpoints will not logged
var IgnoredPathsInLogs = []string{
	"/metrics",
	"/healthz",
}

var cfg *Config

func init() {
	cfg = new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBuffer([]byte(Default))); err != nil {
		log.Fatal("error loading default configs: %v", err)
	}
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.MergeInConfig(); err != nil {
		log.Debug("no config file found. Using defaults and environment variables.")
	}
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatal("invalid config schema: %v", err)
	}
	if err := validator.New().Struct(cfg); err != nil {
		log.Fatal("invalid config: %v", err)
	}
	// notify the developers or user about logs
	log.Debug("we don't log incoming requests to this endpoints: %s", IgnoredPathsInLogs)
	// generate a unique id for this instance
	cfg.App.InstanceID = shortuuid.New()
}

func Set(newConfig *Config) {
	cfg = newConfig
}

func Get() *Config {
	return cfg
}

func (c *Config) IsProduction() bool {
	return c.App.Env == ProductionMode
}

func (c *Config) IsDebugging() bool {
	return c.App.Env == DebuggingMode
}
