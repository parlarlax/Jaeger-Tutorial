package router

import (
	"eaxmple/jaeker/echo/res/handler"
	"errors"

	"github.com/labstack/echo/v4"
)


func Router(e *echo.Echo) error {
	if Prop == nil {
		return errors.New("properties can't be nil must inject properties via Properties func. ")
	}

	handler.Ping(e)
	handler.Pong(e)
	handler.Hello(e)
	
	return nil
}