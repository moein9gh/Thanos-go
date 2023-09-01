package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thanos-go/config"
	"github.com/thanos-go/log"
	"github.com/thoas/go-funk"
)

// RequestsLogger Listens for any errors and log it
func RequestsLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// Start timer
			start := time.Now()
			req := c.Request()
			res := c.Response()

			// Process request
			if err = next(c); err != nil {
				c.Error(err)
			}

			// calculating the request size
			bytesInStr := req.Header.Get(echo.HeaderContentLength)
			if bytesInStr == "" {
				bytesInStr = "0"
			}
			bytesIn, _ := strconv.ParseInt(bytesInStr, 10, 64)

			// Do the logging
			fields := map[string]interface{}{
				"remote_ip":  c.RealIP(),
				"protocol":   req.Proto,
				"method":     req.Method,
				"host":       req.Host,
				"uri":        req.RequestURI,
				"user_agent": req.UserAgent(),
				"latency":    time.Since(start).String(),
				"status":     res.Status,
				"bytes_in":   byteToHumanReadable(bytesIn),
				"bytes_out":  byteToHumanReadable(res.Size),
			}
			// checking if the user is logged in
			if claims, has := c.Get(config.KeyDeviceClaim).(map[string]interface{}); has {
				if id, ok := claims["id"].(float64); ok {
					fields["user_id"] = id
				}
			}

			if err != nil {
				fields["error"] = err.Error()
			}

			// check to skip the ignored endpoints
			if !funk.ContainsString(config.IgnoredPathsInLogs, req.RequestURI) {
				log.Info("incoming request", fields)
			} else if res.Status == http.StatusServiceUnavailable {
				log.Error("incoming request", fields)
			}

			return
		}
	}
}

func byteToHumanReadable(b int64) string {
	if b < 0 {
		return fmt.Sprintf("%v", b)
	}
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
