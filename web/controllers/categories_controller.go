package controllers

import (
			"github.com/nutsdo/taobao-sdk"
			"github.com/kataras/iris"
)

type CategoriesController struct {
	Ctx iris.Context
}

type Category struct {
	name string
	category_id string
	parent_id string
}

func (c *CategoriesController) Get(){

	category := Category{
		name:"女装",
		category_id:"1234",
		parent_id:"1",
	}
	categories := make([]map[string]string,0)
	result := make(map[string]interface{})
	result["status_code"]	= "ok"
	result["data"] = append(categories, taobao.StructToMap(category))
	c.Ctx.JSON(result)
}

func (c * CategoriesController) GetBy(category_id int64){
	category := Category{
		name:"女装",
		category_id:"1234",
		parent_id:"1",
	}

	result := make(map[string]interface{})
	result["status_code"]	= "ok"
	result["data"] = taobao.StructToMap(category)
	c.Ctx.JSON(result)

}