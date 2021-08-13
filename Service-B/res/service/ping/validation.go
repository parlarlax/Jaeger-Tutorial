package ping

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
	TransactionId = "TransactionId"
	RequestDate   = "RequestDate"
	Src           = "source"
	srcIP         = "SourceIP"
	Lang          = "lang"
)

func ValidateRequest(ctx *echo.Context) error {

	return nil
}


func getHeader(headers *http.Header) (auth, tid, reqDate, src, SrcIP, lang string) {
	a := headers.Get(Authorization)
	t := headers.Get(TransactionId)
	r := headers.Get(RequestDate)
	s := headers.Get(Src)
	sr := headers.Get(srcIP)
	l := headers.Get(Lang)
	return a, t, r, s, sr, l
}


func setRespHeader(ctx *echo.Context) {
	timeLayout := "2006-01-02T15:04:05.999"
	responseTime := time.Now().Format(timeLayout)
	(*ctx).Response().Header().Set("ResponseTime", responseTime)
	(*ctx).Response().Header().Set("Content-type", "application/json")
}