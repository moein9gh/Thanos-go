package log

import (
	"log"
	"os"

	"github.com/thanos-go/pkg/logger"

	"github.com/thanos-go/pkg/logger/std_logger"
)

var handler logger.Logger

func init() {
	handler = std_logger.New(os.Stderr, "")
}

func SetAdapter(adapter logger.Logger) {
	handler = adapter
}

func Get() logger.Logger {
	return handler
}

func Std() *log.Logger {
	return handler.Std()
}

func Named(name string) logger.Logger {
	return handler.Named(name)
}

func Debug(msg string, fields ...interface{}) {
	handler.Debug(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	handler.Info(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	handler.Error(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	handler.Fatal(msg, fields...)
}
