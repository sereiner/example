package main

import (
	"apiserver/services/order"
	"github.com/sereiner/parrot/parrot"
)

func main() {

	app := parrot.NewApp(
		parrot.WithPlatName("apiserver"),
		parrot.WithSystemName("apiserver"),
		parrot.WithServerTypes("api"),
		parrot.WithDebug())
	app.Micro("/order/query", order.NewQueryHandler)
	app.Micro("/order/query/bind", order.NewBindHandler)
	app.Start()
}
