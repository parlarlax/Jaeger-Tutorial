package pong

import (
	"context"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/ping"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/tracing"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
)

func Pong(ctx echo.Context) error {
	span := tracing.StartSpanFromRequestPong(*service.Prop.Tracer, ctx.Request())
	defer span.Finish()

	contx := opentracing.ContextWithSpan(context.Background(), span)
	responseB, err := ping.Pong(contx, "localhost:7072")
	if err != nil {
		log.Fatalf("Error occurred: %s\n", err)
	}
	fmt.Println(responseB)

	response := fmt.Sprintf("Pong: Service-A -> %s \n",responseB)
	log.Println(response)
	return ctx.String(http.StatusOK, response)

}