package main

import (
	"rpcserver/services/order"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("rpcserver"),
		parrot.WithSystemName("collector"),
		parrot.WithServerTypes("rpc-api"),
		parrot.WithDebug())

	app.API("/order/query", order.NewQueryHandler)
	app.RPC("/order/bind", order.NewBindHandler)
	app.Start()
}
