package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"blog/controllers/api"
)

func init() {
	//前台
    beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/ajaxarticletable", &controllers.MainController{}, "*:AjaxArticleTable")
	//beego.AutoRouter(&controllers.MainController{})

    //后台
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	//beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.Router("/auth/list", &controllers.AuthController{}, "*:List")
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.ArticleController{})
	beego.AutoRouter(&controllers.UploadController{})

	//api
	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/v1",
				beego.NSRouter("/getarticle", &api.ApiarticleController{}, "get:GetArticle"),
				beego.NSRouter("/gettaglist", &api.ApiarticleController{}, "get:GetTagList"),

			),

		)
	//注册 namespace
	beego.AddNamespace(ns)

	//ns :=
	//	beego.NewNamespace("/frontend",
	//		beego.NSRouter("/index", &frontend.IndexController{}, "get:Index"),
	//	)
	////注册 namespace
	//beego.AddNamespace(ns)

}
