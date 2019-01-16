package main

import (
	"github.com/kataras/iris"
	"github.com/nutsdo/taobao-service/web"
)

func main() {

	app := web.NewApp()

	app.UseGlobal(before)
	app.DoneGlobal(after)

	app.Run(iris.Addr("iris.local:8090"), iris.WithoutServerError(iris.ErrServerClosed))
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Values().Set("info", shareInformation)
	ctx.Next()
}

func after(ctx iris.Context) {
	println("After the indexHandler or contactHandler")
}
