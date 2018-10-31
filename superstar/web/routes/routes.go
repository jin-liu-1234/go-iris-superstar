package routes

import (
	"_iris/superstar/services"
	"_iris/superstar/bootstrap"
	"github.com/kataras/iris/mvc"
	"_iris/superstar/web/controller"
)

func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controller.IndexController))

}
