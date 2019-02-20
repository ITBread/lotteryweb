// webapp project main.go
package main

import "github.com/kataras/iris"

type User struct {
	Username string `json:"name"`
	Password string `json:"age"`
}

func main() {
	app := iris.New()
	//　从　"./views"　目录下加载扩展名是".html" 　的所有模板，
	//　并使用标准的　`html/template`　 包进行解析。
	tmpl := iris.Django("./views", ".html").Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.StaticWeb("/public", "./public") //指定静态文件路径
	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("main.html")
	})

	app.Get("/login", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("login.html")
	})

	app.Get("/index", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		//ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("index.html")
	})

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	app.Get("/user/{id:long}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Post("/user/login", func(ctx iris.Context) {
		c := &User{}
		c.Username = ctx.FormValue("username")
		c.Password = ctx.FormValue("password")
		println("%v   %v", c.Password, c.Username)
		if err := ctx.ReadJSON(c); err != nil {
			ctx.Redirect("/")
		} else {
			ctx.JSON(c)

		}
	})

	//　使用网络地址启动服务
	app.Run(iris.Addr(":8090"))
}
