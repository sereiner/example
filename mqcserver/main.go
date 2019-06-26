package main

import (
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("mqcserver"),
		parrot.WithSystemName("collector"),
		parrot.WithServerTypes("mqc-api"),
		parrot.WithDebug())

	app.Conf.MQC.SetSubConf("server", `
	{
		"proto":"redis",
		"addrs":[
				"192.168.0.111:6379",
				"192.168.0.112:6379"
		],
		"db":1,
		"dial_timeout":10,
		"read_timeout":10,
		"write_timeout":10,
		"pool_size":10
}
`)
	app.Conf.MQC.SetSubConf("queue", `{
		"queues":[
			{
				"queue":"mqc:server",
				"service":"/message/handle"
			}
		]
	}`)
	app.Conf.Plat.SetVarConf("queue", "queue", `
{
	"proto":"redis",
	"addrs":[
			"192.168.0.111:6379",
			"192.168.0.112:6379"
	],
	"db":1,
	"dial_timeout":10,
	"read_timeout":10,
	"write_timeout":10,
	"pool_size":10
}
`)
	app.MQC("/message/handle", msgHandle)
	app.Micro("/message/send", send)
	app.Start()
}
func msgHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------收到消息---------")
	ctx.Log.Info(ctx.Request.GetBody())
	return "success"
}
func send(ctx *context.Context) (r interface{}) {

	queue, err := ctx.GetContainer().GetQueue()
	if err != nil {
		return err
	}
	if err = queue.Push("mqc:server", `{"id":"1001"}`); err != nil {
		return err
	}
	return "success"
}
