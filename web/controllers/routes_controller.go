package controllers

import (
	"github.com/kataras/iris"
	"reflect"
	"fmt"
)

type RoutesController struct {
	Ctx iris.Context
}

func (c *RoutesController) Get(){

	r :=c.Ctx.Application().GetRoutesReadOnly()
	for _,v := range r{
		c.Ctx.HTML(v.Name()+"<br/>")

	}

}

func (c *RoutesController) GetReflect(){

	t := reflect.TypeOf(c)
	fmt.Println(t)

	v := reflect.ValueOf(c)

	fmt.Println(v)
}