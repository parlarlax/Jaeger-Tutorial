package hello

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/log"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/jaegertracing"
	"net/http"
	"time"
)

func Hello(ctx echo.Context) error {
	//span := tracing.StartSpanFromRequest(*service.Prop.Tracer, ctx.Request())
	//defer span.Finish()
	////
	response := `{"status": true, "message":"ระบบดำเนินการสำเร็จ", "data:":"Hello World!"}`
	//span.LogFields(
	//	log.String("event", response),
	//	log.String("event", "soft error"),
	//	log.String("type", "cache timeout"),
	//	log.Int("waited.millis", 1500))
	//span.SetBaggageItem("Test baggage", "baggage")
	//span.SetTag("Test tag", "New Tag")
	//span.SetTag("response", response)
	//span.SetBaggageItem("response", response)
	//span.SetOperationName("Hello")
	//
	//return ctx.String(http.StatusOK, response)

	// Do something before creating the child span
	//time.Sleep(40 * time.Millisecond)
	sp := jaegertracing.CreateChildSpan(ctx, "Child span for additional processing")
	defer sp.Finish()
	sp.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 1500))
	sp.SetBaggageItem("Test baggage", "baggage")
	sp.SetBaggageItem("Test response", response)
	sp.SetTag("Test tag", "New Tag")
	time.Sleep(100 * time.Millisecond)
	return ctx.String(http.StatusOK, "Hello, World!")

}
