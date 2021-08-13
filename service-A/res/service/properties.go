package service

import "eaxmple/jaeker/echo/app"

var Prop *app.Properties

func Properties(ap *app.Properties) {
	Prop = ap
}