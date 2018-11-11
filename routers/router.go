package routers

import (
	"blog/controllers"
	"blog/controllers/api"

	"github.com/astaxie/beego"
)

func init() {

	//前台
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/ajaxarticletable", &controllers.MainController{}, "*:AjaxArticleTable")
	//文章列表
	beego.Router("/farticle", &controllers.MainController{}, "*:Article")
	//文章详情页
	beego.Router("/fdetail", &controllers.MainController{}, "*:Detail")
	beego.Router("/ftimeline", &controllers.MainController{}, "*:Timeline")
	beego.Router("/fabout", &controllers.MainController{}, "*:About")
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
				//hot top列表
				beego.NSRouter("/getranklist", &api.ApiarticleController{}, "get:GetRankList"),

				//留言列表
				beego.NSRouter("/getcomment", &api.ApicommentController{}, "get:GetComment"),
				//留言
				beego.NSRouter("/postcomment", &api.ApicommentController{}, "post:PostComment"),

				//emeiym留言
				beego.NSRouter("/emeiympostcomment", &api.ApiEmeiymCommentController{}, "post:PostComment"),
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
