// webapp project main.go
package main

import (
	"github.com/ITBread/lotteryweb/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析。
	tmpl := iris.Django("./views", ".html").Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.StaticWeb("/public", "./public") //指定静态文件路径
	users := mvc.New(app.Party("/ssq"))  // 一个根路径为 /users 的组
	//users.Register(userService)
	users.Handle(new(controllers.SsqController)) // 定义组的controller

	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件：
		ctx.View("index.html")
	})
	//　使用网络地址启动服务
	app.Run(iris.Addr(":8090"))
}
