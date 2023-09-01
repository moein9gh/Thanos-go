package zap_logger

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/thanos-go/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	LogPath    string `mapstructure:"log-path"`
	LocalTime  bool   `mapstructure:"local-time"`
	MaxSize    int    `mapstructure:"max-size"`
	MaxAge     int    `mapstructure:"max-age"`
	MaxBackups int    `mapstructure:"max-backups"`
	Compress   bool   `mapstructure:"compress"`
}

type ZapAdapter struct {
	handler *zap.Logger
	named   string
}

func New(isProduction bool, config Config, level zapcore.Level) (*ZapAdapter, func()) {

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.ConsoleSeparator = " "

	cores := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			level),
	}

	if isProduction && config.LogPath != "" {
		cores = append(cores,
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(&lumberjack.Logger{
					Filename:   config.LogPath,
					MaxSize:    config.MaxSize,
					MaxAge:     config.MaxAge,
					MaxBackups: config.MaxBackups,
					LocalTime:  config.LocalTime,
					Compress:   config.Compress,
				}),
				zap.InfoLevel))
	}

	internalLogger := zap.New(zapcore.NewTee(cores...))

	return &ZapAdapter{
			handler: internalLogger,
		}, func() {
			_ = internalLogger.Sync()
		}
}

func (za *ZapAdapter) Std() *log.Logger {
	return zap.NewStdLog(za.handler)
}

func (za *ZapAdapter) Get() logger.Logger {
	return za
}

func (za *ZapAdapter) Named(name string) logger.Logger {
	za.named = name
	return za
}

func (za *ZapAdapter) Debug(msg string, fields ...interface{}) {
	msg, fields = za.formatMessageWithValues(msg, fields...)
	za.handler.Named(za.named).Debug(msg, za.normalizeValues(fields)...)
	za.named = ""
}

func (za *ZapAdapter) Info(msg string, fields ...interface{}) {
	msg, fields = za.formatMessageWithValues(msg, fields...)
	za.handler.Named(za.named).Info(msg, za.normalizeValues(fields)...)
	za.named = ""
}

func (za *ZapAdapter) Error(msg string, fields ...interface{}) {
	msg, fields = za.formatMessageWithValues(msg, fields...)
	za.handler.Named(za.named).Error(msg, za.normalizeValues(fields)...)
	za.named = ""
}

func (za *ZapAdapter) Fatal(msg string, fields ...interface{}) {
	msg, fields = za.formatMessageWithValues(msg, fields...)
	za.handler.Named(za.named).Fatal(msg, za.normalizeValues(fields)...)
	za.named = ""
}

func (za *ZapAdapter) formatMessageWithValues(msg string, values ...interface{}) (string, []interface{}) {
	thresholds := strings.Count(msg, "%")
	returnedValues := make([]interface{}, 0)
	if thresholds > 0 {
		if len(values) >= thresholds {
			msg = fmt.Sprintf(msg, values[:thresholds]...)
			if len(values) > thresholds {
				returnedValues = values[thresholds:]
			}
		} else if len(values) > 0 {
			msg = fmt.Sprintf(msg, values...)
		}
	} else if len(values) > 0 {
		returnedValues = values[thresholds:]
	}
	return msg, returnedValues
}

func (za *ZapAdapter) normalizeValues(values []interface{}) []zap.Field {
	fields := make([]zap.Field, 0)
	if len(values) == 1 {
		switch typedValue := values[0].(type) {
		case error:
			fields = append(fields, zap.Error(typedValue))
		case map[string]string:
			for k, v := range typedValue {
				fields = append(fields, zap.String(k, v))
			}
		case map[string]interface{}:
			for k, v := range typedValue {
				fields = append(fields, zap.Any(k, v))
			}
		default:
			fields = append(fields, zap.String("extra", fmt.Sprintf("%v", typedValue)))
		}
	} else if len(values) > 1 {
		fields = append(fields, zap.String("extra", fmt.Sprintf("%v", values[1:])))
	}
	return fields
}
