package ping

import (
	"eaxmple/jaeker/echo/lib/tracing"
	"eaxmple/jaeker/echo/res/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Ping(ctx echo.Context) error {
	span := tracing.StartSpanFromRequestPing(*service.Prop.Tracer, ctx.Request())
	defer span.Finish()

	response := `Service-D`
	log.Println("Ping:",response)
	return ctx.String(http.StatusOK, response)
}



