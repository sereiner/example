package main

import (
	"webserver/services/order"
	"webserver/services/user"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("webserver"),
		parrot.WithSystemName("collector"),
		parrot.WithServerTypes("web"),
		parrot.WithDebug())

	app.Web("/user/login", user.NewLoginHandler)
	app.Web("/order/query", order.NewQueryHandler)
	app.Micro("/order/bind", order.NewBindHandler)

	app.Start()
}
