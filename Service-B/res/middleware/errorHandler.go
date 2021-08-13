package middleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

const (
	DefaultErr      = "SYS_ACC_ERR_0000"
	DefaultErrTH    = `เกิดข้อผิดพลาดภายในระบบ`
	DefaultErrMsgTH = `ระบบไม่สามารถดำเนินการได้`
	DefaultErrEN    = `The internal system error has occurred.`
	DefaultErrMsgEN = `The system cannot operate.`
	OracleNoListener = "ORA-12541: TNS:no listener\n"
	DriverBadConn    = "driver: bad connection"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Errorf("%v",err)
	switch err.Error() {
	case OracleNoListener:
		noListener(&c)
	case DriverBadConn:
		driverBadConnect(&c)
	default:
		if err := errIsOutSideTheControl(&c, err); err != nil {
			return
		}
		c.JSON(http.StatusInternalServerError, err)
	}
}

//Some error that cannot be controlled
func errIsOutSideTheControl(ctx *echo.Context, err error) error {
	c := *ctx
	fool := make(map[string]interface{})
	err_ := json.Unmarshal([]byte(err.Error()), &fool)
	if err_ != nil {
		err__ := buildDefaultErr(c.Request().Header.Get("Lang"))
		c.JSON(http.StatusInternalServerError, err__)
	}
	return err_
}

func noListener(ctx *echo.Context) {
	c := *ctx
	err := buildDefaultErr(c.Request().Header.Get("Lang"))
	c.JSON(http.StatusInternalServerError, err)
}

func driverBadConnect(ctx *echo.Context) {
	c := *ctx
	err := buildDefaultErr(c.Request().Header.Get("Lang"))
	c.JSON(http.StatusInternalServerError, err)
}

func buildDefaultErr(lang string) map[string]interface{} {
	desc := make(map[string]interface{})
	errDetail := make(map[string]string)
	errDetail["code"] = DefaultErr
	errDetails := make([]map[string]string, 0)
	errDetails = append(errDetails, errDetail)
	desc["errors"] = errDetails
	desc["status"] = false
	switch strings.ToUpper(lang) {
	case "EN":
		errDetail["description"] = DefaultErrEN
		desc["message"] = DefaultErrMsgEN
	default:
		errDetail["description"] = DefaultErrTH
		desc["message"] = DefaultErrMsgTH
	}
	return desc
}