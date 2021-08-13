package handler

import (
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/middleware"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service/hello"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service/ping"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service/pong"

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