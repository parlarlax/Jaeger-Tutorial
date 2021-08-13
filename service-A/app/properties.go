package app

import (
	"database/sql"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Properties struct {
	Viper     *viper.Viper
	Oracle    *sql.DB
	Log       *logrus.Logger
	MsgJson   *string
	Tracer    *opentracing.Tracer
}