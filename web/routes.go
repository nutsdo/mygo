package web

import (
	"github.com/kataras/iris/mvc"
	"github.com/nutsdo/taobao-service/web/controllers"
	"github.com/kataras/iris"
	"github.com/nutsdo/taobao-service/web/controllers/tools"
		)


func Routes(app *iris.Application){

	app.UseGlobal(func(ctx iris.Context) {
		sign := ctx.URLParam("sign")
		if sign!="auth" {
			//ctx.JSON("1111111")
		}
		//for _,v := range app.GetRoutes(){
		//	ctx.HTML(v.Name + "</br>")
		//}

		ctx.Next()
	})
	//Enforcer := casbin.NewEnforcer("./middlewares/casbinmodel.conf", "./middlewares/casbinpolicy.csv")
	//Enforcer.LoadModel()
	//Enforcer.LoadPolicy()
	//Enforcer.SavePolicy()
	//casbinMiddleware := cm.New(Enforcer)
	//app.Use(casbinMiddleware.ServeHTTP)
	//mvc 模式
	//商品相关路由
	product :=mvc.New(app.Party("/product"))
	product.Handle(new(controllers.ProductController))

	//分类路由组
	categories := mvc.New(app.Party("/categories"))
	//categories.Router.Use(middleware.BasicAuth)
	categories.Handle(new(controllers.CategoriesController))

	//管理员接口
	admin :=mvc.New(app.Party("/admins"))
	admin.Handle(new(controllers.AdminController))
	//oauth2.0
	auth := mvc.New(app.Party("/auth"))
	auth.Handle(new(controllers.AuthorizationController))


	//一些小工具
	tools_sign := mvc.New(app.Party("/tools"))
	tools_sign.Handle(new(tools.ToolsSignController))


	//路由列表
	routes := mvc.New(app.Party("/routes"))
	routes.Handle(new(controllers.RoutesController))
	routes.Register(app)
}

