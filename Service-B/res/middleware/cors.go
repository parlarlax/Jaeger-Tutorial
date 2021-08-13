package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			Skipper:          nil,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowCredentials: true,
			ExposeHeaders:    nil,
			MaxAge:           0,
		},
	)
}
