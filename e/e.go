package e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thanos-go/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/p3ym4n/re"
)

var isProduction = true

func ProductionStatus(status bool) {
	isProduction = status
}

func LogAndDeliver(err re.Error) (int, map[string]interface{}) {
	Log(err)
	return Deliver(err)
}

func Log(err re.Error) {
	if err.Kind() != re.KindInvalid && err.Kind() != re.KindNotFound {
		log.Error(err.Message(), err.ProcessedMap())
	}
}

func Deliver(err re.Error) (int, map[string]interface{}) {
	internal := err.Internal()
	code := re.HttpCode(err)
	msg := err.Message()

	switch code {
	case http.StatusUnprocessableEntity:
		out := map[string]interface{}{
			"message": "request parameters are not valid",
		}
		if msg != "" {
			out["error"] = parseStringToValidation(msg)
		} else if asserted, ok := internal.(validation.Errors); ok {
			out["error"] = convertErrorToMap(asserted)
		} else {
			out["error"] = parseStringToValidation(internal.Error())
		}
		return code, out
	case http.StatusForbidden:
		if msg == "" {
			msg = internal.Error()
			if isProduction {
				msg = "the action is not allowed"
			}
		}
		return code, map[string]interface{}{
			"message": msg,
		}
	case http.StatusNotFound:
		if msg == "" {
			msg = internal.Error()
			if isProduction {
				msg = "the requested url not found"
			}
		}
		return code, map[string]interface{}{
			"message": msg,
		}
	default:
		out := map[string]interface{}{
			"message": "an error occurred",
		}
		if !isProduction && msg != "" {
			out["error"] = msg
		}
		return code, out
	}
}

func BindingError(err error) (int, map[string]interface{}) {
	if asserted, ok := err.(*echo.HTTPError); ok {
		err = asserted.Internal
	}
	return http.StatusBadRequest, map[string]interface{}{
		"message": "cannot bind the request body",
		"error":   err.Error(),
	}
}

func parseStringToValidation(msg string) interface{} {
	if !strings.ContainsAny(msg, ";") && !strings.ContainsAny(msg, ":") {
		return msg
	}
	list := make(map[string]string)
	for i, item := range strings.Split(msg, "; ") {
		parts := strings.SplitN(item, ": ", 2)
		if len(parts) > 1 {
			list[parts[0]] = strings.Join(parts[1:], ":")
		} else {
			list[fmt.Sprintf("error #%v", i+1)] = item
		}
	}
	return list
}

func convertErrorToMap(es validation.Errors) map[string]interface{} {
	errs := map[string]interface{}{}
	for key, err := range es {
		if ms, ok := err.(json.Marshaler); ok {
			errs[key] = ms
		} else {
			errs[key] = err.Error()
		}
	}
	return errs
}
