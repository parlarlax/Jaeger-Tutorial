package middleware

import (
	ping "eaxmple/jaeker/echo/res/service/ping"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func ValidatePing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ping.ValidateRequest(&c); err != nil {
			log.Errorf("%v",err)
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
