package middlewareCustom

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func EventLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC3339), c.Request().Method, c.Request().URL.Path)
		return next(c)
	}
}
