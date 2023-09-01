package logger

import "log"

type Logger interface {
	Std() *log.Logger
	Get() Logger
	Named(string) Logger
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}
