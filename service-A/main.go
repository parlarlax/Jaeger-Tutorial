package main

import (
	"fmt"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/app"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/di"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/jaegertracing"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/tracing"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/middleware"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/router"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

const(
	thisServiceName = "service-a"
)

var prop *app.Properties

func bootstrap() error {
	//init config
	prop = new(app.Properties)

	//init tracer
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	prop.Tracer = &tracer

	//init Log
	log := logrus.New()
	var loglvl logrus.Level
	rawLogLvl := "info"
	if rawLogLvl == "" {
		rawLogLvl = "debug"
	}
	loglvl, err := logrus.ParseLevel(rawLogLvl)
	if err != nil {
		return err
	}
	log.SetLevel(loglvl)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	prop.Log = log

	di.Inject(prop)
	return nil
}

func serveRestAPI() {
	e := echo.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()
	e.Use(middleware.CORS())
	if err := router.Router(e); err != nil {
		prop.Log.Panicln(err)
	}
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	host := "localhost"
	port := "7071"
	if err := e.Start(fmt.Sprintf("%s:%s", host, port)); err != nil {
		prop.Log.Panicln("Failed to server Rest %v", err)
	}
}


func main() {
	
	err := bootstrap()
	if err != nil {
		fmt.Println(fmt.Errorf("main error: %v", err))
	}
	serveRestAPI()
}







