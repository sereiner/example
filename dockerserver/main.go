package main

import (
	"github.com/sereiner/parrot/context"
	"github.com/sereiner/parrot/parrot"
)

func main() {
	app := parrot.NewApp(
		parrot.WithPlatName("myplat"),   //平台名称
		parrot.WithSystemName("demo"),   //系统名称
		parrot.WithClusterName("test"),  //集群名称
		parrot.WithServerTypes("api"),   //只启动http api 服务
		parrot.WithRegistry("fs://../"), //使用本地文件系统作为注册中心
		parrot.WithDebug())
	app.Micro("/hello", helloWorld)
	app.Start()
}

func helloWorld(ctx *context.Context) (r interface{}) {
	return "hello world"
}
