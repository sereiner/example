package main

import (
	"cronserver/services/order"
	"cronserver/services/user"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("cronserver"),
		parrot.WithSystemName("collector"),
		parrot.WithServerTypes("cron"),
		parrot.WithDebug())

	app.CRON("/user/login", user.NewLoginHandler)
	app.CRON("/order/query", order.NewQueryHandler)
	app.CRON("/order/bind", order.NewBindHandler)
	app.CRON("/order/bind", order.NewBindHandler)
	app.Start()
}
