package ping

import (
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/ping"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/lib/tracing"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"

	"context"
	"github.com/labstack/echo/v4"
)

func Ping(ctx echo.Context) error {
	span := tracing.StartSpanFromRequestPing(*service.Prop.Tracer, ctx.Request())
	defer span.Finish()

	contx := opentracing.ContextWithSpan(context.Background(), span)
	responseB, err := ping.Ping(contx, "localhost:7072")
	if err != nil {
		log.Fatalf("Error occurred: %s\n", err)
	}

	fmt.Println("Receive from: ", responseB)
	//w.Write([]byte(fmt.Sprintf("[ %s -> %s, ", thisServiceName, responseB)))


	responseC, err := ping.Ping(contx, "localhost:7073")
	if err != nil {
		log.Fatalf("Error occurred: %s\n", err)
	}
	fmt.Println("Receive from: ", responseC)
	//w.Write([]byte(fmt.Sprintf("%s -> %s, ", thisServiceName, responseC)))


	responseD, err := ping.Ping(contx, "localhost:7074")
	if err != nil {
		log.Fatalf("Error occurred: %s\n", err)
	}
	fmt.Println("Receive from: ", responseD)
	//w.Write([]byte(fmt.Sprintf("%s -> %s ]\n", thisServiceName, responseD)))

	response := fmt.Sprintf("Ping: Service-A -> %s -> %s -> %s \n",responseB, responseC, responseD)
	log.Println(response)

	
	return ctx.String(http.StatusOK, response)
}


