package config

import (
	"time"

	"github.com/thanos-go/pkg/logger/zap_logger"
)

type Config struct {
	App            AppConfig   `mapstructure:"app"`
	Mysql          MysqlConfig `mapstructure:"mysql"`
	Authentication `mapstructure:"authentication"`
	ZapLogger      zap_logger.Config `mapstructure:"zap-logger"`
	HashID         `mapstructure:"hash-id" `
	Static         Static `mapstructure:"static" validate:"required"`
}

type AppConfig struct {
	Name             string        `mapstructure:"name" validate:"required"`
	Env              string        `mapstructure:"env" validate:"required"`
	Port             int           `mapstructure:"port" validate:"required"`
	GracefulShutdown time.Duration `mapstructure:"graceful-shutdown" validate:"required"`
	CorsEnabled      bool          `mapstructure:"cors-enabled" validate:""`
	RequestTimeout   time.Duration `mapstructure:"request-timeout" validate:"required"`
	LogPath          string        `mapstructure:"log-path" validate:""`
	InstanceID       string        `mapstructure:"" validate:""`
	ScheduleLockTTL  time.Duration `mapstructure:"schedule-lock-ttl" validate:""`
}

type MysqlConfig struct {
	Host            string        `mapstructure:"host" validate:"required"`
	Schema          string        `mapstructure:"schema" validate:"required"`
	Port            string        `mapstructure:"port" validate:"required"`
	Username        string        `mapstructure:"username" validate:"required"`
	Password        string        `mapstructure:"password" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max-open-conns" validate:""`
	MaxIdleConns    int           `mapstructure:"max-idle-conns" validate:""`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime" validate:""`
}

type Authentication struct {
	AccessExpirationInMinute  int    `mapstructure:"access-expiration-in-minute"`
	RefreshExpirationInMinute int    `mapstructure:"refresh-expiration-in-minute"`
	JwtSecret                 string `mapstructure:"jwt-secret"`
}

type HashID struct {
	Salt      string `mapstructure:"thanos-hash-id-salt"`
	MinLength int    `mapstructure:"min-length"`
}

type Static struct {
	StaticFilePath    string `mapstructure:"static-file-path"`
	CharacterFilePath string `mapstructure:"character-file-path"`
	ImageBase         string `mapstructure:"image-base" validate:"required"`
}
