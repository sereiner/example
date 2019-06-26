package main

import (
	"github.com/sereiner/parrot/context"

	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("metricserver"),
		parrot.WithSystemName("metricserver"),
		parrot.WithServerTypes("api"),
	)

	app.Conf.API.SetMainConf(`{"address":":9067"}`)
	app.Conf.API.SetSubConf("metric", `{
	"host":"http://192.168.106.219:8086",     
	"dataBase":"metricserver",
	"cron":"@every 10s",
	"userName":"",
	"password":""
	}	`)

	app.API("/hello", helloWorld)
	app.Start()
}

func helloWorld(ctx *context.Context) (r interface{}) {
	//panic("hello.err")
	return "hello"
}
