package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type HomeController struct {

}

var HomeStaticView = mvc.View{
	Name:"/home.html",
	Data: iris.Map{"Title": "Iyese"},
}

func (c *HomeController) Get() mvc.Result{

	return HomeStaticView

}