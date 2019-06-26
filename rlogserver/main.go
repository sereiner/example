package main

import (
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/parrot"
)

type OrderResult struct {
	OrderNO     string `xml:"qxOrderNo"`
	OrderStatus string `xml:"orderStatus"`
	ErrCode     string `xml:"failedCode"`
	ErrMsg      string `xml:"failedReason"`
}

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("rlogger"),
		parrot.WithSystemName("rlogger"),
		parrot.WithServerTypes("api"),
		parrot.WithRemoteLogger(),
		parrot.WithDebug())

	app.Conf.API.SetMainConf(`{"address":":7892","trace":true}`)

	app.Conf.Plat.SetVarConf("global", "logger", `{
    "level":"All",
		"interval":"1s",
		"layout":{"name":"%name","time":"%datetime","content":"%content","level":"%l","session":"%session"},
    "service":"/parrot/log/save@logging.logging_debug"
}`)
	app.Conf.API.SetSubConf("metric", `{
	"host":"http://192.168.0.185:8086",
	"dataBase":"rlogserver",
	"cron":"@every 10s",
	"userName":"",
	"password":""
	}`)

	app.Micro("/hello", helloWorld)
	app.Start()
}

func helloWorld(ctx *context.Context) (r interface{}) {
	ctx.Response.SetXML()
	return &OrderResult{
		OrderNO:     "QX09099999",
		OrderStatus: "UNDERWAY",
		ErrCode:     "0001",
		ErrMsg:      "success",
	}
}
