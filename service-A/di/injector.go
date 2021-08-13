package di

import (
	"github.com/parlarlax/Jaeger-Tutorial/service-A/app"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/middleware"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/router"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/service"
	"github.com/parlarlax/Jaeger-Tutorial/service-A/res/handler"
)


func Inject(ap *app.Properties) {
	injectRest(ap)
	injectDomain(ap)
	injectAdapter(ap)
}

func injectDomain(ap *app.Properties) {
	// msgCode.Properties(ap)
	// searchDomain.Properties(ap)
}

func injectAdapter(ap *app.Properties) {
	// oracleConnecting.Properties(ap)
	// searchRepo.Properties(ap)
}

func injectRest(ap *app.Properties) {
	router.Properties(ap)
	middleware.Properties(ap)
	handler.Properties(ap)
	service.Properties(ap)
}