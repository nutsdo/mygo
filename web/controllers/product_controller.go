package controllers

import (
						"github.com/kataras/iris"
	"github.com/nutsdo/taobao-sdk"
	"github.com/nutsdo/taobao-sdk/requests/taobaoke"
	"strconv"
)

type ProductController struct {
	Ctx iris.Context
}

type Product struct {
	title string
	price string
	coupon string
}

func (c *ProductController) GetDetailBy(product_id int64){
	//params := c.Ctx.Params();
	//fmt.Println()
	//查询商品信息
	tbClient := taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}
	productId := strconv.FormatInt(product_id,10)
	req := new(taobaoke.GetItemInfoRequest)
	req.SetNumIids(productId)
	respMap, err := tbClient.Invoke(req)
	if err != nil{
		c.Ctx.JSON(err)
	}else {
		c.Ctx.JSON(respMap)
	}
}