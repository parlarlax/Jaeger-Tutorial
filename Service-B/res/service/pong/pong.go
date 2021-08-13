package pong

import (
	"eaxmple/jaeker/echo/lib/tracing"
	"eaxmple/jaeker/echo/res/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Pong(ctx echo.Context) error {
	span := tracing.StartSpanFromRequestPong(*service.Prop.Tracer, ctx.Request())
	defer span.Finish()
	//
	//contx := opentracing.ContextWithSpan(context.Background(), span)
	//responseD, err := ping.Pong(contx, "localhost:7074")
	//if err != nil {
	//	log.Fatalf("Error occurred: %s\n", err)
	//}
	//
	//fmt.Println(responseD)
	//response := fmt.Sprintf("Service-C -> %s", responseD)
	response := `Service-D`
	log.Println("Pong:",response)
	return ctx.String(http.StatusOK, response)
}
