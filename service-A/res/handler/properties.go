package handler

import "github.com/parlarlax/Jaeger-Tutorial/service-A/app"

var Prop *app.Properties

func Properties(ap *app.Properties) {
	Prop = ap
}