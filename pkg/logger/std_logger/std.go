package std_logger

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/thanos-go/pkg/logger"
)

type StdAdapter struct {
	handler *log.Logger
	named   string
}

func New(output io.Writer, prefix string) *StdAdapter {
	return &StdAdapter{
		handler: log.New(output, prefix, 0),
	}
}

func (sa *StdAdapter) Std() *log.Logger {
	return sa.handler
}

func (sa *StdAdapter) Get() logger.Logger {
	return sa
}

func (sa *StdAdapter) Named(name string) logger.Logger {
	sa.named = name
	return sa
}

func (sa *StdAdapter) Debug(msg string, fields ...interface{}) {
	msg = sa.setup("DEBUG", msg)
	msg, fields = sa.formatMessageWithValues(msg, fields)
	fieldsCount := len(fields)
	if fieldsCount == 0 {
		sa.handler.Println(msg)
		return
	} else if fieldsCount == 1 {
		sa.handler.Println(msg, fields[0])
		return
	}
	sa.handler.Println(msg, fields)
}

func (sa *StdAdapter) Info(msg string, fields ...interface{}) {
	msg = sa.setup("INFO", msg)
	msg, fields = sa.formatMessageWithValues(msg, fields)
	fieldsCount := len(fields)
	if fieldsCount == 0 {
		sa.handler.Println(msg)
		return
	} else if fieldsCount == 1 {
		sa.handler.Println(msg, fields[0])
		return
	}
	sa.handler.Println(msg, fields)
}

func (sa *StdAdapter) Error(msg string, fields ...interface{}) {
	msg = sa.setup("ERROR", msg)
	msg, fields = sa.formatMessageWithValues(msg, fields)
	fieldsCount := len(fields)
	if fieldsCount == 0 {
		sa.handler.Println(msg)
		return
	} else if fieldsCount == 1 {
		sa.handler.Println(msg, fields[0])
		return
	}
	sa.handler.Println(msg, fields)
}

func (sa *StdAdapter) Fatal(msg string, fields ...interface{}) {
	msg = sa.setup("FATAL", msg)
	msg, fields = sa.formatMessageWithValues(msg, fields)
	fieldsCount := len(fields)
	if fieldsCount == 0 {
		sa.handler.Println(msg)
		return
	} else if fieldsCount == 1 {
		sa.handler.Println(msg, fields[0])
		return
	}
	sa.handler.Println(msg, fields)
}

func (sa *StdAdapter) formatMessageWithValues(msg string, values []interface{}) (string, []interface{}) {
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

func (sa *StdAdapter) setup(level, msg string) string {
	timeStamp := time.Now().Format(time.RFC3339)
	sa.handler.SetPrefix(fmt.Sprintf("%s %s", timeStamp, level))
	if sa.named != "" {
		msg = fmt.Sprintf("%s %s ", sa.named, msg)
	}
	return msg
}
