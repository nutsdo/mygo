package web

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/nutsdo/taobao-sdk"
	"github.com/nutsdo/taobao-sdk/requests/taobaoke"
	"github.com/nutsdo/taobao-sdk/requests/taobaoke/sc"
	"github.com/nutsdo/taobao-sdk/requests/taobaoke/tools"
	"time"
	)

func NewApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover2.New())
	app.Use(logger.New())

	assetHandler := app.StaticHandler("./public", false, false)

	app.SPA(assetHandler)

	tpl := "public"
	//template := iris.HTML("./web/views/templates/"+tpl, ".html")
	template := iris.HTML("./"+tpl, ".html")
	template.Reload(true)
	app.RegisterView(template)

	//app.Get("/", Home)
	app.Get("/version", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"description": "淘宝API", "version": "1.0.0", "author": "lvdingtao"})
	})

	app.Get("/taobao/getItem", GetTBItem)
	app.Get("/taobao/createTlj", CreateTaoLiJin)
	app.Get("/taobao/couponBrandRecommendSc", CouponBrandRecommendSc)
	app.Post("/taobao/CreateTbPwd", CreateTbPwd)
	app.Get("/taobao/GetPrivilege", GetPrivilege)

	//路由包
	Routes(app)

	return app
}

func Home(ctx iris.Context) {

	ctx.View("index.html")

}

func GetPrivilege(ctx iris.Context) {
	tbClient := taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}

	req := new(tools.GetPrivilegeRequest)
	//mm_99317136_20476650_69686669
	//mm_48807125_22220223_81486002
	//req.SetCampaignType("")
	req.SetValue("session", "620220972095def3acec70d2ebd5a323134de28bdbe092a1866605434")
	req.SetAdzoneId("81486002")
	req.SetSiteId("22220223")
	req.SetItemId("42861838038")
	respMap, err := tbClient.Invoke(req)
	ctx.JSON(err)
	ctx.JSON(respMap)
}

func CreateTbPwd(ctx iris.Context) {

	params := ctx.FormValues()
	fmt.Println(params)
	tbClient := taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}

	req := new(taobaoke.TbkCreateTbpwdRequest)
	couponUrl := ctx.PostValueDefault("couponUrl", "")
	if couponUrl == "" {
		//iris.StatusUnprocessableEntity
		return
	}
	//https://uland.taobao.com/taolijin/edetail?eh=ROCROBp1ohmZuQF0XRz0iAXoB+DaBK5LQS0Flu/fbSp4QsdWMikAalrisGmre1Id0BFAqRODu10xGp9u6ozyABON/nWC2a7ro1K0x8s2HfhDcRDefEt7KIfDCBqp97nfdfvIQHIqZlQ//HbTHjt8hgBexaqSu3eb3R2WU4nWMhjUXf8v8BviREWaGx75npJgxE+c6rQ1iD5tHi6XRm2oI3aaJd7fEF7aIw9CHS9QH/3RhhlBT4XK1tdUpkX+TQUC1Cygw4KSB77n0Pj5JCgADWUFFgz3RBVH&union_lens=lensId:0b15677f_0b3c_1667fb348cc_0739;traffic_flag=lm
	req.SetUrl(couponUrl)

	req.SetText("测试淘礼金转链")
	respMap, err := tbClient.Invoke(req)
	ctx.JSON(err)
	ctx.JSON(respMap)

}

func GetTBItem(ctx iris.Context) {
	fmt.Println(ctx.URLParams())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	var tbClient = taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}
	req := new(taobaoke.TbkGetItemRequest)
	req.SetFields("num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick")
	req.SetQ("女装")
	respMap, err := tbClient.Invoke(req)
	ctx.JSON(err)
	ctx.JSON(respMap)

}

func CreateTaoLiJin(ctx iris.Context) {
	var tbClient = taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}
	req := new(sc.TbkCreateTaoLiJinScRequest)
	//mm_99317136_20476650_69686669
	//req.SetCampaignType("")
	req.SetValue("session", "6201f160e6ad57d86c0ce40ceg50066ed99a3324d3cd443133167095545")
	req.SetAdzoneId("43798100236")
	req.SetSiteId("33434789")
	req.SetItemId("561540532406")
	req.SetSendEndTime("2018-10-20 00:00:00")
	req.SetTotalNum("2")
	req.SetName("淘礼金试用")
	req.SetUserToTalWinNumLimit("1")
	req.SetSecuritySwitch("true")
	req.SetPerFace("1")
	req.SetSendStartTime("2018-10-16 23:59:59")
	respMap, err := tbClient.Invoke(req)
	ctx.JSON(err)
	ctx.JSON(respMap)
}

func CouponBrandRecommendSc(ctx iris.Context) {
	var tbClient = taobao.TaobaoApp{
		"23617975",
		"3f722bb9f2f90ef770da88d08b363f6b",
	}
	req := new(sc.CouponBrandRecommendScRequest)
	req.SetValue("session", "6201d11f6acd25b3064ZZ7e9cf801886112daf7a98819de2257973620")
	req.SetAdzoneId("69686669")
	req.SetSiteId("20476650")
	respMap, err := tbClient.Invoke(req)
	//ctx.JSON(err)
	fmt.Println(err)
	ctx.JSON(respMap)
}
