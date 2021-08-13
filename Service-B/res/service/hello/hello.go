package hello

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Hello(ctx echo.Context) error {
	//span := tracing.StartSpanFromRequest(*service.Prop.Tracer, ctx.Request())
	//defer span.Finish()

	response := `{"status": true, "message":"ระบบดำเนินการสำเร็จ", "data:":"Hello World!"}`
	log.Println(response)
	return ctx.String(http.StatusOK, response)
}