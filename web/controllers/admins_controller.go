package controllers

import (
	"github.com/kataras/iris"
	"github.com/nutsdo/taobao-service/models"
			)

type AdminController struct {
	Ctx iris.Context
}


func (c *AdminController) Get(){

	su := new(models.SystemUsers)

	var id int64

	id = 1
	su = su.GetById(id)


	c.Ctx.JSON(su)

	c.Ctx.HTML("获取管理员列表")
}

