package http

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

var App *iris.Application
var IrisInterface IIris

type IIris interface {
	Init()
	GetApp() *iris.Application
	RegisterPostRouter(path string, handler func(ctx iris.Context), callback func(ctx iris.Context))
	RegisterGetRouter(path string, handler func(ctx iris.Context), callback func(ctx iris.Context))
	Cors(ctx iris.Context)
}
type Iris struct {
}

func NewIrIrisInterface() IIris {
	App = iris.Default()
	IrisInterface = &Iris{}
	return IrisInterface
}

func (i *Iris) Init() {
	app := i.GetApp()
	app.Use(i.Cors)
	app.Use(recover.New())
	app.Run(iris.Addr(":8999"), iris.WithoutServerError(iris.ErrServerClosed))

}

func (i *Iris) GetApp() *iris.Application {
	return App
}

func (i *Iris) RegisterPostRouter(path string, handler func(ctx iris.Context), callback func(ctx iris.Context)) {

	app := i.GetApp()
	if handler == nil {
		app.Post(path, callback)
	} else {
		app.Post(path, handler, callback)
	}
	app.Logger().Info("[注册路由]版本: v1 " + "请求类型:" + " POST" + " 请求路径: /v1" + path)
}
func (i *Iris) RegisterGetRouter(path string, handler func(ctx iris.Context), callback func(ctx iris.Context)) {
	app := i.GetApp()
	if handler == nil {
		app.Get(path, callback)
	} else {
		app.Get(path, handler, callback)
	}
	app.Logger().Info("[注册路由]版本: v1 " + "请求类型:" + " GET" + " 请求路径: " + path)
}
func (i *Iris) Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
