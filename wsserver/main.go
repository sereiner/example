package main

import (
	"wsserver/services/order"
	"wsserver/services/user"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("ws-server"),
		parrot.WithSystemName("collector"),
		parrot.WithServerTypes("api-ws"),
		parrot.WithDebug())

	app.WS("/auth/login", user.NewLoginHandler)
	app.WS("/query", order.NewQueryHandler)
	app.WS("/recv", order.NewBindHandler)
	app.Conf.WS.SetSubConf("auth", `
		{
			"jwt": {
				"exclude": ["/auth/login"],
				"expireAt": 36000,
				"mode": "HS512",
				"name": "__jwt__",
				"redirect":"/auth/login",
				"secret": "12345678"
			}
		}
		`)
	app.Start()
}
