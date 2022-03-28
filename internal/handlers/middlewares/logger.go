package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Logger is log request
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		if err != nil {
			return err
		}

		logs := logrus.Fields{
			"host":         c.Hostname(),
			"method":       c.Method(),
			"path":         c.OriginalURL(),
			"language":     c.Locals(context.LangKey),
			"ip":           c.IP(),
			"user_agent":   c.Get(fiber.HeaderUserAgent),
			"body_size":    fmt.Sprintf("%.5f MB", float64(bytes.NewReader(c.Request().Body()).Len())/1024.00/1024.00),
			"status_code":  fmt.Sprintf("%d %s", c.Response().StatusCode(), fasthttp.StatusMessage(c.Response().StatusCode())),
			"process_time": time.Since(start),
		}

		employeeID := context.New(c).GetEmployeeID()
		if employeeID != "" {
			logs["employee_id"] = employeeID
		}

		parameters := c.Locals(context.ParametersKey)
		if parameters != nil {
			b, _ := json.Marshal(parameters)
			for _, f := range []string{"Password"} {
				if res := gjson.GetBytes(b, f); res.Exists() {
					b, _ = sjson.SetBytes(b, f, "**********")
				}
			}
			logs["parameters"] = string(b)
		} else {
			logs["parameters"] = "{}"
		}

		return nil
	}
}
