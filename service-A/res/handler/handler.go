package handler

import (
	"eaxmple/jaeker/echo/res/middleware"
	"eaxmple/jaeker/echo/res/service/hello"
	"eaxmple/jaeker/echo/res/service/ping"
	"eaxmple/jaeker/echo/res/service/pong"

	"github.com/labstack/echo/v4"
)



func Ping(e *echo.Echo)  {
	e.GET("/ping",
		ping.Ping,
		middleware.ValidatePing,
	)
}

func Pong(e *echo.Echo) {
	e.GET("/pong",
		pong.Pong,
		middleware.ValidatePong,
	)
}

func Hello(e *echo.Echo) {
	e.GET("/hello",
		hello.Hello,
		middleware.ValidateHello,
	)
}